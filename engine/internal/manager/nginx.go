package manager

import (
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
