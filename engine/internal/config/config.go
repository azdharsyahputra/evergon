package config

type Config struct {
	ServerAddr      string `json:"server_addr"`
	Workspace       string `json:"workspace"`
	PHPExecutable   string `json:"php_executable"`
	NginxExecutable string `json:"nginx_executable"`
	TemplateDir     string `json:"template_dir"`
	NginxVHostDir   string `json:"nginx_vhost_dir"`
}

func Load() Config {
	return Config{
		ServerAddr:      "127.0.0.1:9090",
		Workspace:       "C:/Dev",
		PHPExecutable:   "C:/Evergon/php_versions/php81/php-fpm.exe",
		NginxExecutable: "C:/Evergon/nginx/nginx.exe",
		TemplateDir:     "C:/Evergon/nginx_template",
		NginxVHostDir:   "C:/Evergon/nginx/conf/vhosts",
	}
}
