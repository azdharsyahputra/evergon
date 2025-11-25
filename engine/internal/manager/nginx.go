package manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/process"
	"evergon/engine/internal/util/resolver"
)

func StartNginx(res *resolver.Resolver) error {
	return process.Start(res.NginxBinary(), "-c", res.NginxConf())
}

func StopNginx(res *resolver.Resolver) error {
	return process.Stop(res.NginxBinary())
}

func ReloadNginx(res *resolver.Resolver) error {
	return process.Run(res.NginxBinary(), "-s", "reload", "-c", res.NginxConf())
}

func CreateVHost(domain, root, phpPort string, res *resolver.Resolver) error {
	tmplPath := filepath.Join(res.TemplateDir(), "vhost.conf")
	raw, err := os.ReadFile(tmplPath)
	if err != nil {
		return fmt.Errorf("failed to read vhost template: %v", err)
	}

	content := string(raw)

	phpBlock := `
    location ~ \.php$ {
        include fastcgi_params;
        fastcgi_pass 127.0.0.1:` + phpPort + `;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
    }`

	content = strings.ReplaceAll(content, "{{SERVER_NAME}}", domain)
	content = strings.ReplaceAll(content, "{{ROOT_PATH}}", root)
	content = strings.ReplaceAll(content, "{{PHP_BLOCK}}", phpBlock)

	if _, err := os.Stat(res.VHostDir()); os.IsNotExist(err) {
		os.MkdirAll(res.VHostDir(), 0755)
	}

	output := res.VHostFile(domain)

	if err := os.WriteFile(output, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write vhost file: %v", err)
	}

	return ReloadNginx(res)
}
