package api

import (
	"log"
	"net/http"

	"evergon/engine/internal/config"
	"evergon/engine/internal/util/resolver"
)

func StartServer(cfg config.Config, res *resolver.Resolver) {
	mux := http.NewServeMux()

	RegisterRoutes(mux, res)

	log.Println("[API] Listening on", cfg.ServerAddr)
	if err := http.ListenAndServe(cfg.ServerAddr, mux); err != nil {
		log.Fatalf("[API ERROR] %v", err)
	}
}
