package pathresolver

import (
	"os"
	"path/filepath"
	"runtime"
)

type Config struct {
	// Root direktori utama Evergon, misal: C:\Evergon atau /opt/evergon
	Root string

	// Subfolder relatif dari Root
	PhpVersionsDir   string // contoh: "php_versions"
	NginxTemplateDir string // contoh: "nginx_template"
	SitesDir         string // contoh: "sites" atau "projects"
}

// Resolver bertugas ngerakit path final.
type Resolver struct {
	cfg      Config
	isWin    bool
	rootPath string
}

// NewResolver menerima config mentah.
// Kalau Root kosong, dia fallback ke direktori executable.
func NewResolver(cfg Config) (*Resolver, error) {
	isWin := runtime.GOOS == "windows"

	root := cfg.Root
	if root == "" {
		exe, err := os.Executable()
		if err != nil {
			return nil, err
		}
		// misal: /path/to/evergon/engine/cmd/evergon-engine/evergon-engine
		// lu bisa adjust naik berapa level sesuai struktur real.
		root = filepath.Dir(filepath.Dir(filepath.Dir(exe)))
	}

	return &Resolver{
		cfg:      cfg,
		isWin:    isWin,
		rootPath: root,
	}, nil
}

// Root mengembalikan path root absolut Evergon.
func (r *Resolver) Root() string {
	return r.rootPath
}

// PhpBaseDir: C:\Evergon\php_versions
func (r *Resolver) PhpBaseDir() string {
	return filepath.Join(r.rootPath, r.cfg.PhpVersionsDir)
}

// NginxTemplateBaseDir: C:\Evergon\nginx_template
func (r *Resolver) NginxTemplateBaseDir() string {
	return filepath.Join(r.rootPath, r.cfg.NginxTemplateDir)
}

// SitesBaseDir: C:\Evergon\sites
func (r *Resolver) SitesBaseDir() string {
	return filepath.Join(r.rootPath, r.cfg.SitesDir)
}

// PHPBinary mengembalikan path binary PHP berdasarkan "tag versi" Evergon.
// Misal: "81" -> C:\Evergon\php_versions\php81\php.exe (Windows)
//
//	"81" -> /opt/evergon/php_versions/php81/bin/php (Linux)
func (r *Resolver) PHPBinary(version string) string {
	base := r.PhpBaseDir()

	if r.isWin {
		// C:\Evergon\php_versions\php81\php.exe
		return filepath.Join(base, "php"+version, "php.exe")
	}

	// /opt/evergon/php_versions/php81/bin/php
	return filepath.Join(base, "php"+version, "bin", "php")
}

// PHPIni mengembalikan path file php.ini untuk versi tertentu.
func (r *Resolver) PHPIni(version string) string {
	base := r.PhpBaseDir()
	if r.isWin {
		return filepath.Join(base, "php"+version, "php.ini")
	}
	return filepath.Join(base, "php"+version, "php.ini")
}

// NginxBinary mengembalikan path nginx.
// Kalau lu ship nginx sendiri, letakkan di Root/nginx/nginx(.exe)
func (r *Resolver) NginxBinary() string {
	name := "nginx"
	if r.isWin {
		name += ".exe"
	}
	return filepath.Join(r.rootPath, "nginx", name)
}

// NginxConf: file vhost untuk site tertentu.
// Misal: C:\Evergon\nginx_template\evergon.local.conf
func (r *Resolver) NginxConf(site string) string {
	return filepath.Join(r.NginxTemplateBaseDir(), site+".conf")
}

// SiteRoot: direktori root project/site.
// Misal: C:\Evergon\sites\evergon-panel
func (r *Resolver) SiteRoot(site string) string {
	return filepath.Join(r.SitesBaseDir(), site)
}
