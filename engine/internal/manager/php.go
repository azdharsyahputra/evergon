package manager

import (
	"evergon/engine/internal/config"
	"fmt"
	"os/exec"
)

var phpCmd *exec.Cmd

func StartPHP(root string) error {
	cfg := config.Load()

	if cfg.PHPMode != "builtin" {
		// FPM mode (Linux default service)
		return nil
	}

	// Kalau sudah running, jangan start dua kali
	if phpCmd != nil {
		return fmt.Errorf("PHP already running")
	}

	// Start PHP built-in server
	phpCmd = exec.Command(cfg.PHPExecutable, "-S", "127.0.0.1:9000", "-t", root)
	phpCmd.Stdout = nil
	phpCmd.Stderr = nil

	if err := phpCmd.Start(); err != nil {
		phpCmd = nil
		return fmt.Errorf("PHP built-in start failed: %v", err)
	}

	return nil
}

func StopPHP() error {
	cfg := config.Load()

	if cfg.PHPMode != "builtin" {
		// FPM mode: nothing to stop
		return nil
	}

	// Built-in process kill
	if phpCmd != nil {
		err := phpCmd.Process.Kill()
		phpCmd = nil
		return err
	}

	// fallback kill
	return exec.Command("pkill", "-f", "php -S").Run()
}
