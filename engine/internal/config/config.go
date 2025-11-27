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

	// === NEW: GLOBAL PHP CONFIG ===
	GlobalPHPPort   int    `json:"global_php_port"`
	GlobalPublicDir string `json:"global_public_dir"`
}

func Load() Config {
	isWin := runtime.GOOS == "windows"

	var root string

	// 1. ENV override
	if envRoot := os.Getenv("EVERGON_ROOT"); envRoot != "" {
		root = envRoot
	} else if cwd, err := os.Getwd(); err == nil {
		// 2. go run mode (development)
		dir := filepath.Dir(cwd)
		if _, err := os.Stat(dir); err == nil {
			root = dir
		}
	} else if exe, err := os.Executable(); err == nil {
		// 3. binary mode (portable)
		dir := filepath.Dir(filepath.Dir(filepath.Dir(exe)))
		if _, err := os.Stat(dir); err == nil {
			root = dir
		}
	} else {
		// 4. fallback
		home, _ := os.UserHomeDir()
		if isWin {
			root = filepath.Join(home, "Evergon")
		} else {
			root = filepath.Join(home, "evergon")
		}
	}

	// === Default paths ===
	workspace := filepath.Join(root, "workspace")
	publicDir := filepath.Join(workspace, "public")

	phpVer := "81"
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
		phpExec = filepath.Join(phpBase, "bin", "php")
		nginxExec = filepath.Join(nginxBase, "portable", "sbin", "nginx")
		nginxConf = filepath.Join(nginxBase, "portable", "conf", "nginx.conf")
		nginxVhosts = filepath.Join(nginxBase, "portable", "conf", "vhosts")
	}

	// return final config
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

		// === GLOBAL PHP SETTINGS ===
		GlobalPHPPort:   8000,
		GlobalPublicDir: publicDir,
	}
}
