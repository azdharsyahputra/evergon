package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
)

// ===========================
// NGINX PROCESS CONTROL
// ===========================

func StartNginx() error {
	cfg := config.Load()
	return process.Start(cfg.NginxExecutable, "-c", cfg.NginxConf)
}

func StopNginx() error {
	// Linux: killall nginx
	// Windows: taskkill nginx.exe
	return process.Stop("nginx")
}

func ReloadNginx() error {
	cfg := config.Load()
	return process.Run(cfg.NginxExecutable, "-s", "reload", "-c", cfg.NginxConf)
}

// ===========================
// VHOST GENERATOR (HYBRID)
// ===========================

func CreateVHost(domain, root, phpPort string) error {
	cfg := config.Load()

	// 1. Load template
	tmplPath := filepath.Join(cfg.TemplateDir, "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read vhost template: %v", err)
	}

	content := string(raw)

	// 2. Generate PHP BLOCK based on OS mode
	var phpBlock string

	if cfg.PHPMode == "fpm" {
		// Linux PHP-FPM
		phpBlock = `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass unix:` + cfg.FPMSocket + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
	} else {
		// Windows Built-in PHP Server
		phpBlock = `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:` + phpPort + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
	}

	// 3. Replace variables
	content = strings.ReplaceAll(content, "{{SERVER_NAME}}", domain)
	content = strings.ReplaceAll(content, "{{ROOT_PATH}}", root)
	content = strings.ReplaceAll(content, "{{PHP_BLOCK}}", phpBlock)

	// 4. Ensure vhost dir exists
	if _, err := os.Stat(cfg.NginxVHostDir); os.IsNotExist(err) {
		os.MkdirAll(cfg.NginxVHostDir, 0755)
	}

	// 5. Write vhost file
	output := filepath.Join(cfg.NginxVHostDir, domain+".conf")

	err = os.WriteFile(output, []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write vhost file: %v", err)
	}

	// 6. Reload nginx
	return ReloadNginx()
}
