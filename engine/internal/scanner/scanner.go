package scanner

import (
	"os"
	"path/filepath"

	"evergon/engine/internal/util/resolver"
)

type Project struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func detectType(path string) string {
	if fileExists(filepath.Join(path, "artisan")) {
		return "laravel"
	}
	if fileExists(filepath.Join(path, "system")) {
		return "ci4"
	}
	if fileExists(filepath.Join(path, "wp-config.php")) {
		return "wordpress"
	}
	return "unknown"
}

func Scan(res *resolver.Resolver) []Project {
	root := filepath.Join(res.Workspace(), "www")
	result := []Project{}

	entries, err := os.ReadDir(root)
	if err != nil {
		return result
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		projectPath := filepath.Join(root, entry.Name())
		projectType := detectType(projectPath)

		result = append(result, Project{
			Name: entry.Name(),
			Path: projectPath,
			Type: projectType,
		})
	}

	return result
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
