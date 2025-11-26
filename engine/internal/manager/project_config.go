package manager

import (
	"encoding/json"
	"os"
	"path/filepath"

	"evergon/engine/internal/config"
)

type ProjectConfig struct {
	PHPVersion string `json:"php_version"`
	PHPPort    string `json:"php_port"`
}

func configPath(projectName string) string {
	cfg := config.Load()
	return filepath.Join(cfg.Workspace, "www", projectName, ".evergon.json")
}

func LoadProjectConfig(projectName string) (*ProjectConfig, error) {
	path := configPath(projectName)

	raw, err := os.ReadFile(path)
	if err != nil {
		// default config if not exist
		return &ProjectConfig{
			PHPVersion: "",
			PHPPort:    "",
		}, nil
	}

	var cfg ProjectConfig
	err = json.Unmarshal(raw, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SaveProjectConfig(projectName string, cfg *ProjectConfig) error {
	path := configPath(projectName)

	raw, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, raw, 0644)
}
