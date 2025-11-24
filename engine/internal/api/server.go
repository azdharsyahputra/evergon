package api

import (
	"log"
	"net/http"

	"evergon/engine/internal/config"
)

func StartServer(cfg config.Config) {
	mux := http.NewServeMux()

	// register routes
	RegisterRoutes(mux)

	log.Println("[API] Listening on", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, mux); err != nil {
		log.Fatalf("[API ERROR] %v", err)
	}
}
