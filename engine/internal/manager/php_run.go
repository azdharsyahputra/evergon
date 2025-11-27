package manager

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"

	"evergon/engine/internal/config"
)

var phpProcesses = map[string]*exec.Cmd{}

func findFreePort() (string, error) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		return "", err
	}
	defer l.Close()
	return strconv.Itoa(l.Addr().(*net.TCPAddr).Port), nil
}

func isPortInUse(port string) bool {
	conn, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return true
	}
	conn.Close()
	return false
}

func clearStaleProcess(projectName string, port string) {
	if !isPortInUse(port) {
		delete(phpProcesses, projectName)
	}
}

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
		port, err := findFreePort()
		if err != nil {
			return "", err
		}
		pcfg.PHPPort = port
		SaveProjectConfig(projectName, pcfg)
	} else {
		if isPortInUse(pcfg.PHPPort) {
			clearStaleProcess(projectName, pcfg.PHPPort)
			if isPortInUse(pcfg.PHPPort) {
				return "", fmt.Errorf("port %s already in use", pcfg.PHPPort)
			}
		}
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

	if cmd, ok := phpProcesses[projectName]; ok {
		if cmd.ProcessState == nil {
			return pcfg.PHPPort, nil
		}
		delete(phpProcesses, projectName)
	}

	cmd := exec.Command(phpExec, "-S", "127.0.0.1:"+pcfg.PHPPort, "-t", projectPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = nil
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start PHP: %v", err)
	}

	phpProcesses[projectName] = cmd

	pidPath := filepath.Join(projectPath, ".evergon.pid")
	os.WriteFile(pidPath, []byte(strconv.Itoa(cmd.Process.Pid)), 0644)

	return pcfg.PHPPort, nil
}

func StopProjectPHP(projectName string) error {
	cfg := config.Load()
	projectPath := filepath.Join(cfg.Workspace, "www", projectName)
	pidPath := filepath.Join(projectPath, ".evergon.pid")

	if cmd, ok := phpProcesses[projectName]; ok {
		cmd.Process.Kill()
		delete(phpProcesses, projectName)
		os.Remove(pidPath)
		return nil
	}

	raw, err := os.ReadFile(pidPath)
	if err == nil {
		pid, _ := strconv.Atoi(string(raw))
		process, err := os.FindProcess(pid)
		if err == nil {
			process.Kill()
		}
		os.Remove(pidPath)
	}

	return nil
}

func IsProjectRunning(projectName string) bool {
	cfg := config.Load()
	projectPath := filepath.Join(cfg.Workspace, "www", projectName)
	pidPath := filepath.Join(projectPath, ".evergon.pid")

	raw, err := os.ReadFile(pidPath)
	if err != nil {
		return false
	}

	pid, _ := strconv.Atoi(string(raw))
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}

	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}

	pcfg, _ := LoadProjectConfig(projectName)
	return isPortInUse(pcfg.PHPPort)
}

func RestartProjectPHP(projectName string) (string, error) {
	StopProjectPHP(projectName)
	return StartProjectPHP(projectName)
}
func IsProjectActuallyRunning(project string) bool {
	cfg, _ := LoadProjectConfig(project)
	if cfg.PHPPort == "" {
		return false
	}

	conn, err := net.Listen("tcp", ":"+cfg.PHPPort)
	if err != nil {
		return true // port really in use
	}
	conn.Close()
	return false
}
