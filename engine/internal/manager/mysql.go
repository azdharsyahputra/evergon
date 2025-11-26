package manager

import (
	"evergon/engine/internal/process"
	"evergon/engine/internal/util/resolver"
)

func StartMySQL(path string, res *resolver.Resolver) error {
	return process.Start(path)
}

func StopMySQL(res *resolver.Resolver) error {
	return process.Stop("mysqld.exe")
}
