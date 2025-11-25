package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/process"
	"evergon/engine/internal/util/resolver"
)

func phpBlockFor(projectName string) string {
	cfg, _ := LoadProjectConfig(projectName)

	if cfg.PHPPort != "" {
		return `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:` + cfg.PHPPort + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
	}

	return `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:9000;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`
}

func CreateVHost(projectName string, res *resolver.Resolver) (string, error) {
	domain := projectName + ".local"
	root := res.ProjectRoot(projectName)

	tmplPath := filepath.Join(res.TemplateDir(), "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return "", fmt.Errorf("template missing: %v", err)
	}

	content := string(raw)
	content = strings.ReplaceAll(content, "{{SERVER_NAME}}", domain)
	content = strings.ReplaceAll(content, "{{ROOT_PATH}}", root)
	content = strings.ReplaceAll(content, "{{PHP_BLOCK}}", phpBlockFor(projectName))

	output := filepath.Join(res.NginxVHostDir(), domain+".conf")

	if err := os.WriteFile(output, []byte(content), 0644); err != nil {
		return "", err
	}

	ReloadVHost(res)
	return domain, nil
}

func ReloadVHost(res *resolver.Resolver) error {
	return process.Run(res.NginxExecutable(), "-s", "reload")
}

func ListVHosts(res *resolver.Resolver) []string {
	dir := res.NginxVHostDir()

	entries, err := os.ReadDir(dir)
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
	return ReloadVHost(res)
}

func UpdateVHost(projectName string, res *resolver.Resolver) error {
	_, err := CreateVHost(projectName, res)
	return err
}
