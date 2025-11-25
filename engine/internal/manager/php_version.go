package manager

import (
	"os"
	"path/filepath"
	"strings"

	"evergon/engine/internal/config"
)

type PHPVersion struct {
	Version string `json:"version"`
	Path    string `json:"path"`
}

func DetectPHPVersions() []PHPVersion {
	cfg := config.Load()
	base := filepath.Join(cfg.RootDir, "php_versions")

	entries, err := os.ReadDir(base)
	if err != nil {
		return []PHPVersion{}
	}

	list := []PHPVersion{}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		name := entry.Name() // ex: php81
		ver := strings.TrimPrefix(name, "php")

		var execPath string

		// Linux / Mac
		binPath := filepath.Join(base, name, "bin", "php")
		if fileExists(binPath) {
			execPath = binPath
		}

		// Windows
		winPath := filepath.Join(base, name, "php-cgi.exe")
		if fileExists(winPath) {
			execPath = winPath
		}

		if execPath != "" {
			list = append(list, PHPVersion{
				Version: ver,
				Path:    execPath,
			})
		}
	}

	return list
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
