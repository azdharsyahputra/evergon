package manager

import (
	"fmt"
	"os/exec"

	"evergon/engine/internal/config"
)

func StartPHP(root string) error {
	cfg := config.Load()

	// Linux → using FPM → nothing to start
	if cfg.PHPMode == "fpm" {
		return nil
	}

	// Windows → built-in server
	cmd := exec.Command(cfg.PHPExecutable, "-S", "127.0.0.1:9000", "-t", root)
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start PHP built-in server: %v", err)
	}

	return nil
}

func StopPHP() error {
	cfg := config.Load()

	// Linux → nothing to stop
	if cfg.PHPMode == "fpm" {
		return nil
	}

	// Windows → kill php-cgi / php.exe
	return exec.Command("taskkill", "/IM", "php-cgi.exe", "/F").Run()
}
