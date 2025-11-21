package process

import (
	"os"
	"os/exec"
	"runtime"
)

func Start(path string, args ...string) error {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Start()
}

func Run(path string, args ...string) error {
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Stop process by executable name
func Stop(exe string) error {
	// Windows
	if runtime.GOOS == "windows" {
		return exec.Command("taskkill", "/IM", exe, "/F").Run()
	}

	// Linux / MacOS
	return exec.Command("pkill", "-f", exe).Run()
}
