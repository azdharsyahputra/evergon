package config

import (
	"log"
	"runtime"
)

type Config struct {
	ServerAddr      string `json:"server_addr"`
	Workspace       string `json:"workspace"`
	PHPExecutable   string `json:"php_executable"`
	PHPMode         string `json:"php_mode"` // "fpm" or "builtin"
	FPMSocket       string `json:"fpm_socket"`
	NginxExecutable string `json:"nginx_executable"`
	NginxConf       string `json:"nginx_conf"`
	TemplateDir     string `json:"template_dir"`
	NginxVHostDir   string `json:"nginx_vhost_dir"`
}

func Load() Config {

	// --- WINDOWS MODE (PHP Built-in Server) ---
	if runtime.GOOS == "windows" {
		log.Println("[CONFIG] Windows mode (PHP Built-in Server)")

		return Config{
			ServerAddr:      "127.0.0.1:9090",
			Workspace:       "C:/Evergon/workspace",
			PHPExecutable:   "C:/Evergon/php_versions/php81/php-cgi.exe",
			PHPMode:         "builtin",
			FPMSocket:       "",
			NginxExecutable: "C:/Evergon/nginx/nginx.exe",
			NginxConf:       "C:/Evergon/nginx/conf/nginx.conf",
			TemplateDir:     "C:/Evergon/nginx_template",
			NginxVHostDir:   "C:/Evergon/nginx/conf/vhosts",
		}
	}

	// --- LINUX MODE (PHP-FPM) ---
	log.Println("[CONFIG] Linux mode (PHP-FPM)")

	return Config{
		ServerAddr:      "127.0.0.1:9090",
		Workspace:       "/home/azdhar/evergon/workspace",
		PHPExecutable:   "/usr/sbin/php-fpm8.1", // not used in FPM mode but kept for future
		PHPMode:         "fpm",
		FPMSocket:       "/run/php/php8.1-fpm.sock",
		NginxExecutable: "/home/azdhar/evergon/nginx/portable/sbin/nginx",
		NginxConf:       "/home/azdhar/evergon/nginx/portable/conf/nginx.conf",
		TemplateDir:     "/home/azdhar/evergon/nginx_template",
		NginxVHostDir:   "/home/azdhar/evergon/nginx/portable/conf/vhosts",
	}
}
