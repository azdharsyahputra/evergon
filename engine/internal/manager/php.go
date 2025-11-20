package manager

import (
	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
	"fmt"
)

func StartPHP() error {
	cfg := config.Load()
	phpPath := cfg.PHPExecutable

	if phpPath == "" {
		return fmt.Errorf("PHP path not set in config")
	}

	return process.Start(phpPath)
}

func StopPHP() error {
	return process.Stop("php-fpm.exe")
}
