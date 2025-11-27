package manager

import (
	"fmt"
	"os/exec"
	"strconv"

	"evergon/engine/internal/process"
	"evergon/engine/internal/util/resolver"
)

var phpCmd *exec.Cmd

func StartPHP(root string, port int, res *resolver.Resolver) error {
	if phpCmd != nil {
		return fmt.Errorf("PHP already running")
	}

	binary := res.PHPBinaryFor(GetActivePHPVersion(res))
	addr := "127.0.0.1:" + strconv.Itoa(port)

	phpCmd = exec.Command(binary, "-S", addr, "-t", root)
	phpCmd.Stdout = nil
	phpCmd.Stderr = nil

	if err := phpCmd.Start(); err != nil {
		phpCmd = nil
		return fmt.Errorf("PHP start failed: %v", err)
	}

	return nil
}

func StopPHP(res *resolver.Resolver) error {
	if phpCmd != nil {
		err := phpCmd.Process.Kill()
		phpCmd = nil
		return err
	}

	return process.Stop(res.ActivePHPBinaryName())
}

func IsPHPRunning(res *resolver.Resolver) bool {
	return process.IsRunning(res.ActivePHPBinaryName())
}
