package manager

import (
	"fmt"
	"os/exec"

	"evergon/engine/internal/util/resolver"
)

var phpCmd *exec.Cmd

func StartPHP(root string, res *resolver.Resolver) error {
	if phpCmd != nil {
		return fmt.Errorf("PHP already running")
	}

	phpCmd = exec.Command(res.PHPBinary(), "-S", "127.0.0.1:9000", "-t", root)
	phpCmd.Stdout = nil
	phpCmd.Stderr = nil

	if err := phpCmd.Start(); err != nil {
		phpCmd = nil
		return fmt.Errorf("PHP built-in start failed: %v", err)
	}

	return nil
}

func StopPHP(res *resolver.Resolver) error {
	if phpCmd != nil {
		err := phpCmd.Process.Kill()
		phpCmd = nil
		return err
	}

	return exec.Command("pkill", "-f", res.PHPBinary()).Run()
}
