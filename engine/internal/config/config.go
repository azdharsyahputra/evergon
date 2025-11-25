package config

import (
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	ServerAddr      string `json:"server_addr"`
	RootDir         string `json:"root_dir"`
	Workspace       string `json:"workspace"`
	PHPVersion      string `json:"php_version"`
	PHPExecutable   string `json:"php_executable"`
	PHPMode         string `json:"php_mode"`
	FPMSocket       string `json:"fpm_socket"`
	NginxExecutable string `json:"nginx_executable"`
	NginxConf       string `json:"nginx_conf"`
	TemplateDir     string `json:"template_dir"`
	NginxVHostDir   string `json:"nginx_vhost_dir"`
}

func Load() Config {
	isWin := runtime.GOOS == "windows"

	var root string
	if isWin {
		root = `C:/Evergon`
	} else {
		cwd, _ := os.Getwd()
		root = filepath.Dir(cwd)
	}

	phpVer := "81"
	workspace := filepath.Join(root, "workspace")
	phpBase := filepath.Join(root, "php_versions", "php"+phpVer)
	nginxBase := filepath.Join(root, "nginx")

	var phpExec string
	var nginxExec string
	var nginxConf string
	var nginxVhosts string
	templateDir := filepath.Join(root, "nginx_template")

	if isWin {
		phpExec = filepath.Join(phpBase, "php-cgi.exe")
		nginxExec = filepath.Join(nginxBase, "nginx.exe")
		nginxConf = filepath.Join(nginxBase, "conf", "nginx.conf")
		nginxVhosts = filepath.Join(nginxBase, "conf", "vhosts")
	} else {
		phpExec = "/usr/bin/php"
		nginxExec = filepath.Join(nginxBase, "portable", "sbin", "nginx")
		nginxConf = filepath.Join(nginxBase, "portable", "conf", "nginx.conf")
		nginxVhosts = filepath.Join(nginxBase, "portable", "conf", "vhosts")
	}

	return Config{
		ServerAddr:      "127.0.0.1:9090",
		RootDir:         root,
		Workspace:       workspace,
		PHPVersion:      phpVer,
		PHPExecutable:   phpExec,
		PHPMode:         "builtin",
		FPMSocket:       "",
		NginxExecutable: nginxExec,
		NginxConf:       nginxConf,
		TemplateDir:     templateDir,
		NginxVHostDir:   nginxVhosts,
	}
}
