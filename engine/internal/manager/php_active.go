package manager

import "evergon/engine/internal/util/resolver"

var defaultPHPVersion = "81"

func GetActivePHPVersion(res *resolver.Resolver) string {
	cfg, err := LoadGlobalPHPConfig(res)
	if err != nil || cfg.PHPVersion == "" {
		return defaultPHPVersion
	}
	return cfg.PHPVersion
}
