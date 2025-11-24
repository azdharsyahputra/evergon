package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/process"
	"evergon/engine/internal/scanner"
)

// CORS
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

func RegisterRoutes(mux *http.ServeMux) {

	// -------------------------------------
	// HEALTH
	// -------------------------------------
	mux.HandleFunc("/health", withCORS(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	}))

	// -------------------------------------
	// STATUS
	// -------------------------------------
	mux.HandleFunc("/php/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if process.IsRunning("php -S") {
			fmt.Fprint(w, "running")
		} else {
			fmt.Fprint(w, "stopped")
		}
	}))

	mux.HandleFunc("/nginx/status", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if process.IsRunning("portable/sbin/nginx") {
			fmt.Fprint(w, "running")
		} else {
			fmt.Fprint(w, "stopped")
		}
	}))

	// -------------------------------------
	// PHP BUILT-IN CONTROL
	// -------------------------------------
	mux.HandleFunc("/php/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		root := r.URL.Query().Get("root")
		if root == "" {
			http.Error(w, "root required", 400)
			return
		}

		if err := manager.StartPHP(root); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("PHP started"))
	}))

	mux.HandleFunc("/php/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StopPHP(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("PHP stopped"))
	}))

	// -------------------------------------
	// NGINX CONTROL
	// -------------------------------------
	mux.HandleFunc("/nginx/start", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StartNginx(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx started"))
	}))

	mux.HandleFunc("/nginx/stop", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.StopNginx(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx stopped"))
	}))

	mux.HandleFunc("/nginx/reload", withCORS(func(w http.ResponseWriter, r *http.Request) {
		if err := manager.ReloadNginx(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte("Nginx reloaded"))
	}))

	// -------------------------------------
	// PROJECT SCAN
	// -------------------------------------
	mux.HandleFunc("/projects", withCORS(func(w http.ResponseWriter, r *http.Request) {
		list := scanner.Scan()
		json.NewEncoder(w).Encode(list)
	}))
}
