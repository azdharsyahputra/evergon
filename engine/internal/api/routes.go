package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/process"
	"evergon/engine/internal/scanner"
	"evergon/engine/internal/util/pid"
	"evergon/engine/internal/util/resolver"
)

func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(200)
			return
		}
		fn(w, r)
	}
}

func RegisterRoutes(mux *http.ServeMux, res *resolver.Resolver) {

	RegisterSSE(mux, res)

	mux.HandleFunc("/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if pid.Exists(res.EnginePIDFile()) {
			fmt.Fprint(w, "running")
		} else {
			fmt.Fprint(w, "stopped")
		}
	}))

	mux.HandleFunc("/php/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if process.IsRunning(res.PHPBinary()) {
			fmt.Fprint(w, "running")
		} else {
			fmt.Fprint(w, "stopped")
		}
	}))

	mux.HandleFunc("/nginx/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if process.IsRunning(res.NginxBinary()) {
			fmt.Fprint(w, "running")
		} else {
			fmt.Fprint(w, "stopped")
		}
	}))

	mux.HandleFunc("/php/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		root := r.URL.Query().Get("root")
		if root == "" {
			http.Error(w, "root required", 400)
			return
		}

		if err := manager.StartPHP(res.WorkspaceWWW(), 9000, res); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("PHP started"))
	}))

	mux.HandleFunc("/php/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StopPHP(res); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("PHP stopped"))
	}))

	mux.HandleFunc("/nginx/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StartNginx(res); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx started"))
	}))

	mux.HandleFunc("/nginx/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StopNginx(res); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx stopped"))
	}))

	mux.HandleFunc("/nginx/reload", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.ReloadNginx(res); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx reloaded"))
	}))

	mux.HandleFunc("/projects", withCORS(func(w http.ResponseWriter, r *http.Request) {
		list := scanner.Scan(res)
		json.NewEncoder(w).Encode(list)
	}))

	mux.HandleFunc("/vhost/create", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		if project == "" {
			http.Error(w, "project required", 400)
			return
		}

		// Ensure project exists
		root := filepath.Join(res.WorkspaceWWW(), project)
		if _, err := os.Stat(root); os.IsNotExist(err) {
			http.Error(w, "project not found", 404)
			return
		}

		domain, err := manager.CreateVHost(project, res)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte(domain))
	}))

	mux.HandleFunc("/vhost/list", withCORS(func(w http.ResponseWriter, r *http.Request) {
		list := manager.ListVHosts(res)
		json.NewEncoder(w).Encode(list)
	}))

	mux.HandleFunc("/vhost/remove", withCORS(func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		if domain == "" {
			http.Error(w, "domain required", 400)
			return
		}

		err := manager.RemoveVHost(domain, res)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("removed"))
	}))

	mux.HandleFunc("/vhost/update", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		if project == "" {
			http.Error(w, "project required", 400)
			return
		}

		err := manager.UpdateVHost(project, res)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write([]byte("updated"))
	}))

	mux.HandleFunc("/php/versions", withCORS(func(w http.ResponseWriter, r *http.Request) {
		versions := manager.DetectPHPVersions()
		json.NewEncoder(w).Encode(versions)
	}))

	mux.HandleFunc("/php/project/get", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		cfg, err := manager.LoadProjectConfig(project)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		json.NewEncoder(w).Encode(cfg)
	}))

	mux.HandleFunc("/php/project/set", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		version := r.URL.Query().Get("version")
		port := r.URL.Query().Get("port")

		if project == "" {
			http.Error(w, "project required", 400)
			return
		}

		cfg, _ := manager.LoadProjectConfig(project)

		oldPort := cfg.PHPPort

		if version != "" {
			cfg.PHPVersion = version
		}

		if port != "" {
			cfg.PHPPort = port
		}

		running := manager.IsProjectActuallyRunning(project)

		if running {
			manager.StopProjectPHP(project)
		}

		manager.SaveProjectConfig(project, cfg)

		// only restart if port or version changed
		if running && (oldPort != cfg.PHPPort || version != "") {
			manager.StartProjectPHP(project)
		}

		w.Write([]byte("ok"))
	}))

	mux.HandleFunc("/php/project/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		port, err := manager.StartProjectPHP(project)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "running",
			"project": project,
			"port":    port,
			"url":     "http://127.0.0.1:" + port,
		})
	}))

	mux.HandleFunc("/php/project/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")
		err := manager.StopProjectPHP(project)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("stopped"))
	}))

	mux.HandleFunc("/php/project/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")

		json.NewEncoder(w).Encode(map[string]interface{}{
			"project": project,
			"running": manager.IsProjectRunning(project),
		})
	}))

	mux.HandleFunc("/php/project/restart", withCORS(func(w http.ResponseWriter, r *http.Request) {
		project := r.URL.Query().Get("project")

		port, err := manager.RestartProjectPHP(project)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		json.NewEncoder(w).Encode(map[string]interface{}{
			"status":  "running",
			"project": project,
			"port":    port,
		})
	}))

	mux.HandleFunc("/port/check", withCORS(func(w http.ResponseWriter, r *http.Request) {
		port := r.URL.Query().Get("port")
		if port == "" {
			http.Error(w, "port required", 400)
			return
		}

		available := manager.IsPortAvailable(port)

		json.NewEncoder(w).Encode(map[string]bool{
			"available": available,
		})
	}))

	mux.HandleFunc("/port/suggest", withCORS(func(w http.ResponseWriter, r *http.Request) {
		port := manager.FindAvailablePort()

		json.NewEncoder(w).Encode(map[string]string{
			"port": port,
		})
	}))

	mux.HandleFunc("/php/version/current", withCORS(func(w http.ResponseWriter, r *http.Request) {
		cfgGlobal, _ := manager.LoadGlobalPHPConfig(res)
		json.NewEncoder(w).Encode(cfgGlobal)
	}))

	mux.HandleFunc("/php/version/list", withCORS(func(w http.ResponseWriter, r *http.Request) {
		list := manager.DetectPHPVersions()
		json.NewEncoder(w).Encode(list)
	}))

	mux.HandleFunc("/php/version/set", withCORS(func(w http.ResponseWriter, r *http.Request) {
		version := r.URL.Query().Get("version")
		if version == "" {
			http.Error(w, "version required", 400)
			return
		}

		// validate exist
		valid := false
		for _, v := range manager.DetectPHPVersions() {
			if v.Version == version {
				valid = true
				break
			}
		}
		if !valid {
			http.Error(w, "php version not installed", 400)
			return
		}

		cfgGlobal, _ := manager.LoadGlobalPHPConfig(res)
		cfgGlobal.PHPVersion = version

		// fallback default port kalau belum ada di config
		if cfgGlobal.PHPPort == 0 {
			cfgGlobal.PHPPort = 8000
		}

		manager.SaveGlobalPHPConfig(res, cfgGlobal)

		// restart kalau running
		if manager.IsGlobalPHPRunning() {
			manager.StopGlobalPHP(res)
			manager.StartGlobalPHP(
				res.WorkspacePublic(),
				cfgGlobal.PHPPort,
				cfgGlobal.PHPVersion,
				res,
			)
		}

		fmt.Fprint(w, "ok")
	}))

	mux.HandleFunc("/php/global/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		cfgGlobal, _ := manager.LoadGlobalPHPConfig(res)

		// default port kalau config lama
		if cfgGlobal.PHPPort == 0 {
			cfgGlobal.PHPPort = 8000
		}

		err := manager.StartGlobalPHP(
			res.WorkspacePublic(),
			cfgGlobal.PHPPort,
			cfgGlobal.PHPVersion,
			res,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprint(w, "started")
	}))
	// STOP global PHP server
	mux.HandleFunc("/php/global/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		manager.StopGlobalPHP(res)
		fmt.Fprint(w, "stopped")
	}))

	mux.HandleFunc("/php/global/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		cfgGlobal, _ := manager.LoadGlobalPHPConfig(res)

		if cfgGlobal.PHPPort == 0 {
			cfgGlobal.PHPPort = 8000
		}

		json.NewEncoder(w).Encode(map[string]any{
			"running": manager.IsGlobalPHPRunning(),
			"port":    cfgGlobal.PHPPort,
			"version": cfgGlobal.PHPVersion,
		})
	}))

}
