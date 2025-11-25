package manager

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"evergon/engine/internal/config"
)

var phpProcesses = map[string]*exec.Cmd{}

func StartProjectPHP(projectName string) (string, error) {
	cfg := config.Load()
	projectPath := filepath.Join(cfg.Workspace, "www", projectName)

	versions := DetectPHPVersions()
	pcfg, _ := LoadProjectConfig(projectName)

	if pcfg.PHPVersion == "" {
		if len(versions) == 0 {
			return "", fmt.Errorf("no PHP versions available")
		}
		pcfg.PHPVersion = versions[0].Version
	}

	if pcfg.PHPPort == "" {
		pcfg.PHPPort = GeneratePort()
		SaveProjectConfig(projectName, pcfg)
	}

	var phpExec string
	for _, v := range versions {
		if v.Version == pcfg.PHPVersion {
			phpExec = v.Path
			break
		}
	}

	if phpExec == "" {
		return "", fmt.Errorf("PHP version %s not found", pcfg.PHPVersion)
	}

	if _, ok := phpProcesses[projectName]; ok {
		return pcfg.PHPPort, nil
	}

	cmd := exec.Command(phpExec, "-S", "127.0.0.1:"+pcfg.PHPPort, "-t", projectPath)
	cmd.Stdout = nil
	cmd.Stderr = nil

	err := cmd.Start()
	if err != nil {
		return "", fmt.Errorf("failed to start PHP: %v", err)
	}

	phpProcesses[projectName] = cmd
	return pcfg.PHPPort, nil
}

func StopProjectPHP(projectName string) error {
	if cmd, ok := phpProcesses[projectName]; ok {
		err := cmd.Process.Kill()
		delete(phpProcesses, projectName)
		return err
	}
	return nil
}
