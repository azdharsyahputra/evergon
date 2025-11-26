package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"evergon/engine/internal/api"
	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
	"evergon/engine/internal/util/pid"
	"evergon/engine/internal/util/resolver"
)

func main() {
	cfg := config.Load()
	res := resolver.New(cfg)
	res.EnsureDirs()
	pidFile := res.EnginePIDFile()

	log.Println("[Evergon Engine] Initializing...")

	if pid.Exists(pidFile) {
		existingPID, err := pid.Read(pidFile)
		if err == nil && process.IsRunningPID(existingPID) {
			log.Println("[Evergon Engine] Already running")
			return
		}
		pid.Remove(pidFile)
		log.Println("[Evergon Engine] Stale PID cleaned")
	}

	pid.Write(pidFile, os.Getpid())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		pid.Remove(pidFile)
		os.Exit(0)
	}()

	log.Println("[Evergon Engine] Starting on", cfg.ServerAddr)

	api.StartServer(cfg, res)
}
