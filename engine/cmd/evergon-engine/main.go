package main

import (
	"log"

	"evergon/engine/internal/api"
	"evergon/engine/internal/config"
	"evergon/engine/internal/util/resolver"
)

func main() {
	cfg := config.Load()
	res := resolver.New(cfg)

	log.Println("[Evergon Engine] Starting on", cfg.ServerAddr)
	api.StartServer(cfg, res)
}
