package scanner

import (
	"os"
	"path/filepath"
)

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func Scan() []Project {
	root := "C:/Dev" // nanti ambil dari config
	result := []Project{}

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil || !info.IsDir() {
			return nil
		}

		// Laravel
		if fileExists(filepath.Join(path, "artisan")) {
			result = append(result, Project{filepath.Base(path), path, "laravel"})
		}

		// CI4
		if fileExists(filepath.Join(path, "system")) {
			result = append(result, Project{filepath.Base(path), path, "ci4"})
		}

		// WordPress
		if fileExists(filepath.Join(path, "wp-config.php")) {
			result = append(result, Project{filepath.Base(path), path, "wordpress"})
		}

		return nil
	})

	return result
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
