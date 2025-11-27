package manager

import (
	"encoding/json"
	"os"

	"evergon/engine/internal/util/resolver"
)

type GlobalPHPConfig struct {
	PHPVersion string `json:"php_version"`
	PHPPort    int    `json:"php_port"`
}

func LoadGlobalPHPConfig(res *resolver.Resolver) (GlobalPHPConfig, error) {
	path := res.GlobalPHPConfigFile()

	raw, err := os.ReadFile(path)
	if err != nil {
		// default config
		return GlobalPHPConfig{PHPVersion: "81"}, nil
	}

	var cfg GlobalPHPConfig
	_ = json.Unmarshal(raw, &cfg)

	if cfg.PHPVersion == "" {
		cfg.PHPVersion = "81"
	}

	return cfg, nil
}

func SaveGlobalPHPConfig(res *resolver.Resolver, cfg GlobalPHPConfig) error {
	path := res.GlobalPHPConfigFile()

	b, _ := json.MarshalIndent(cfg, "", "  ")
	return os.WriteFile(path, b, 0644)
}
func GetResolvedGlobalPHPVersion(res *resolver.Resolver) string {
	cfg, err := LoadGlobalPHPConfig(res)
	if err != nil || cfg.PHPVersion == "" {
		return res.PHPVersion() // fallback default dari config
	}
	return cfg.PHPVersion
}
