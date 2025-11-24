package process

import (
	"os"
	"os/exec"
	"runtime"
)

// Start a process without waiting
func Start(path string, args ...string) error {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

// Run and wait until completed
func Run(path string, args ...string) error {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Kill by process name (exe)
func Stop(exe string) error {
	if runtime.GOOS == "windows" {
		return exec.Command("taskkill", "/IM", exe, "/F").Run()
	}
	return exec.Command("pkill", "-f", exe).Run()
}

// Check if process is running
func IsRunning(name string) bool {
	cmd := exec.Command("pgrep", "-f", name)
	err := cmd.Run()
	return err == nil // exit code 0 = FOUND
}
