package main

import (
	"evergon/engine/internal/api"
	"evergon/engine/internal/config"
	"log"
)

func main() {
	cfg := config.Load()

	log.Println("[Evergon Engine] Starting on", cfg.ServerAddr)
	api.StartServer(cfg)
}
