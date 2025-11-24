package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
)

// ===================================
// NGINX CONTROL
// ===================================

func StartNginx() error {
	cfg := config.Load()
	return process.Start(cfg.NginxExecutable, "-c", cfg.NginxConf)
}

func StopNginx() error {
	// stop portable nginx
	return process.Stop("portable/sbin/nginx")
}

func ReloadNginx() error {
	cfg := config.Load()
	return process.Run(cfg.NginxExecutable, "-s", "reload", "-c", cfg.NginxConf)
}

// ===================================
// VHOST GENERATOR
// ===================================

func CreateVHost(domain, root, phpPort string) error {
	cfg := config.Load()

	tmplPath := filepath.Join(cfg.TemplateDir, "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read vhost template: %v", err)
	}

	content := string(raw)

	var phpBlock string

	if cfg.PHPMode == "fpm" {
		phpBlock = `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass unix:` + cfg.FPMSocket + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
	} else {
		phpBlock = `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:` + phpPort + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
	}

	content = strings.ReplaceAll(content, "{{SERVER_NAME}}", domain)
	content = strings.ReplaceAll(content, "{{ROOT_PATH}}", root)
	content = strings.ReplaceAll(content, "{{PHP_BLOCK}}", phpBlock)

	if _, err := os.Stat(cfg.NginxVHostDir); os.IsNotExist(err) {
		os.MkdirAll(cfg.NginxVHostDir, 0755)
	}

	output := filepath.Join(cfg.NginxVHostDir, domain+".conf")

	if err := os.WriteFile(output, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write vhost file: %v", err)
	}

	return ReloadNginx()
}
