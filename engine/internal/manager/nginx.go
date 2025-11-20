package manager

import (
	"fmt"
	"os"
	"path/filepath"

	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
	"evergon/engine/internal/util"
)

func StartNginx() error {
	cfg := config.Load()
	return process.Start(cfg.NginxExecutable)
}

func StopNginx() error {
	return process.Stop("nginx.exe")
}

func ReloadNginx() error {
	return process.Run("nginx.exe", "-s", "reload")
}

func CreateVHost(domain string, root string, phpPort string) error {
	cfg := config.Load()

	// Load template
	tmplPath := filepath.Join(cfg.TemplateDir, "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read template: %v", err)
	}

	data := map[string]string{
		"SERVER_NAME": domain,
		"ROOT_PATH":   root,
		"PHP_PORT":    phpPort,
	}

	rendered := util.ReplaceAll(string(raw), data)

	// Save to nginx conf.d
	output := filepath.Join(cfg.NginxVHostDir, domain+".conf")
	if err := util.WriteFile(output, rendered); err != nil {
		return err
	}

	return ReloadNginx()
}
