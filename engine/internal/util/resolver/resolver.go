package resolver

import (
	"path/filepath"

	"evergon/engine/internal/config"
)

type Resolver struct {
	cfg config.Config
}

func New(cfg config.Config) *Resolver {
	return &Resolver{cfg: cfg}
}

func (r *Resolver) Root() string {
	return r.cfg.RootDir
}

func (r *Resolver) Workspace() string {
	return r.cfg.Workspace
}

func (r *Resolver) PHPBinary() string {
	return r.cfg.PHPExecutable
}

func (r *Resolver) PHPVersion() string {
	return r.cfg.PHPVersion
}

func (r *Resolver) NginxBinary() string {
	return r.cfg.NginxExecutable
}

func (r *Resolver) NginxConf() string {
	return r.cfg.NginxConf
}

func (r *Resolver) TemplateDir() string {
	return r.cfg.TemplateDir
}

func (r *Resolver) VHostDir() string {
	return r.cfg.NginxVHostDir
}

func (r *Resolver) VHostFile(name string) string {
	return filepath.Join(r.cfg.NginxVHostDir, name+".conf")
}

func (r *Resolver) SiteRoot(name string) string {
	return filepath.Join(r.cfg.Workspace, name)
}
func (r *Resolver) EnginePIDFile() string {
	return filepath.Join(r.cfg.RootDir, "engine.pid")
}
func (r *Resolver) NginxVHostDir() string {
	return r.cfg.NginxVHostDir
}

func (r *Resolver) WorkspaceWWW() string {
	return filepath.Join(r.cfg.Workspace, "www")
}

func (r *Resolver) LogDir() string {
	return filepath.Join(r.cfg.RootDir, "logs")
}

func (r *Resolver) PHPVersionDir() string {
	return filepath.Join(r.cfg.RootDir, "php_versions")
}
