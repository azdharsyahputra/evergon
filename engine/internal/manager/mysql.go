package manager

import "evergon/engine/internal/process"

func StartMySQL(path string) error {
	return process.Start(path)
}

func StopMySQL() error {
	return process.Stop("mysqld.exe")
}
