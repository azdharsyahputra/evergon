package main

import (
	"log"

	"evergon/engine/internal/api"
	"evergon/engine/internal/config"
)

func main() {
	cfg := config.Load()

	log.Println("[Evergon Engine] Starting on", cfg.ServerAddr)
	api.StartServer(cfg)
}
