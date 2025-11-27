package manager

import (
	"encoding/json"
	"os"

	"evergon/engine/internal/config"
	"evergon/engine/internal/util/resolver"
)

type GlobalPHPConfig struct {
	PHPVersion string `json:"php_version"`
}

func LoadGlobalPHPConfig() (GlobalPHPConfig, error) {
	res := resolver.New(config.Load())
	path := res.GlobalPHPConfigFile()

	raw, err := os.ReadFile(path)
	if err != nil {
		return GlobalPHPConfig{PHPVersion: "81"}, nil
	}

	var cfg GlobalPHPConfig
	json.Unmarshal(raw, &cfg)
	return cfg, nil
}

func SaveGlobalPHPConfig(cfg GlobalPHPConfig) error {
	res := resolver.New(config.Load())
	path := res.GlobalPHPConfigFile()

	b, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, b, 0644)
}
