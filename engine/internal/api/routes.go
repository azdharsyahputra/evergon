package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/process"
	"evergon/engine/internal/scanner"
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
	mux.HandleFunc("/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
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

		if err := manager.StartPHP(root, res); err != nil {
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
}
