package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/scanner"
)

func RegisterRoutes(mux *http.ServeMux) {

	// Health Check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	// ============================
	//        PHP CONTROLS
	// ============================
	mux.HandleFunc("/php/start", func(w http.ResponseWriter, r *http.Request) {
		root := r.URL.Query().Get("root")
		if root == "" {
			http.Error(w, "root required", 400)
			return
		}

		err := manager.StartPHP(root)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprint(w, "PHP started")
	})

	mux.HandleFunc("/php/stop", func(w http.ResponseWriter, r *http.Request) {
		err := manager.StopPHP()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, "PHP stopped")
	})

	// ============================
	//        NGINX CONTROLS
	// ============================
	mux.HandleFunc("/nginx/start", func(w http.ResponseWriter, r *http.Request) {
		err := manager.StartNginx()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, "Nginx started")
	})

	mux.HandleFunc("/nginx/stop", func(w http.ResponseWriter, r *http.Request) {
		err := manager.StopNginx()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, "Nginx stopped")
	})

	mux.HandleFunc("/nginx/reload", func(w http.ResponseWriter, r *http.Request) {
		err := manager.ReloadNginx()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		fmt.Fprint(w, "Nginx reloaded")
	})

	// ============================
	//        PROJECT SCAN
	// ============================
	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		list := scanner.Scan()
		json.NewEncoder(w).Encode(list)
	})

	// ============================
	//     VHOST GENERATION
	// ============================
	mux.HandleFunc("/vhost/create", func(w http.ResponseWriter, r *http.Request) {
		domain := r.URL.Query().Get("domain")
		root := r.URL.Query().Get("root")
		phpPort := r.URL.Query().Get("php_port")

		if domain == "" || root == "" || phpPort == "" {
			http.Error(w, "domain, root, php_port required", 400)
			return
		}

		err := manager.CreateVHost(domain, root, phpPort)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		fmt.Fprint(w, "Vhost created & nginx reloaded")
	})
}
