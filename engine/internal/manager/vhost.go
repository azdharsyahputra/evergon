package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/config"
	"evergon/engine/internal/process"
	"evergon/engine/internal/util/resolver"
)

func CreateVHost(domain, root, phpPort string, res *resolver.Resolver) error {
	cfg := config.Load()

	tmplPath := filepath.Join(cfg.TemplateDir, "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("template missing: %v", err)
	}

	content := string(raw)

	phpBlock := `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:9000;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`

	content = strings.ReplaceAll(content, "{{SERVER_NAME}}", domain)
	content = strings.ReplaceAll(content, "{{ROOT_PATH}}", root)
	content = strings.ReplaceAll(content, "{{PHP_BLOCK}}", phpBlock)

	output := filepath.Join(cfg.NginxVHostDir, domain+".conf")
	return os.WriteFile(output, []byte(content), 0644)
}

func ReloadVHost() error {
	cfg := config.Load()
	return process.Run(cfg.NginxExecutable, "-s", "reload")
}

func ListVHosts(res *resolver.Resolver) []string {
	cfg := config.Load()
	entries, err := os.ReadDir(cfg.NginxVHostDir)
	if err != nil {
		return []string{}
	}
	out := []string{}
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".conf") {
			out = append(out, e.Name())
		}
	}
	return out
}

func RemoveVHost(domain string, res *resolver.Resolver) error {
	path := filepath.Join(res.NginxVHostDir(), domain+".conf")
	os.Remove(path)
	return ReloadVHost()
}
