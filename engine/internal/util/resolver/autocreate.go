package resolver

import (
	"log"
	"os"
)

func ensure(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
		log.Println("[AutoCreate]", path, "created")
	} else {
		log.Println("[AutoCreate]", path, "exists")
	}
}

func (r *Resolver) EnsureDirs() {
	ensure(r.cfg.Workspace)
	ensure(r.WorkspaceWWW())
	ensure(r.NginxVHostDir())
	ensure(r.LogDir())
	ensure(r.PHPVersionDir())
}
