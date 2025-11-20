package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"evergon/engine/internal/manager"
	"evergon/engine/internal/scanner"
)

func RegisterRoutes(mux *http.ServeMux) {

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	// PHP
	mux.HandleFunc("/php/start", func(w http.ResponseWriter, r *http.Request) {
		err := manager.StartPHP()
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

	// Scanner API
	mux.HandleFunc("/projects", func(w http.ResponseWriter, r *http.Request) {
		list := scanner.Scan()
		json.NewEncoder(w).Encode(list)
	})

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
