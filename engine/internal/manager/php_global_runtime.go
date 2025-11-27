package manager

import (
	"evergon/engine/internal/util/resolver"
	"fmt"
	"net"
	"os/exec"
	"strconv"
	"time"
)

var globalPHPCmd *exec.Cmd
var globalPHPPort int

func StartGlobalPHP(root string, port int, version string, res *resolver.Resolver) error {

	if globalPHPCmd != nil {
		return fmt.Errorf("global PHP already running")
	}

	phpBin := res.PHPBinaryFor(version)
	addr := "127.0.0.1:" + strconv.Itoa(port)

	cmd := exec.Command(phpBin, "-S", addr, "-t", root)

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start global PHP: %v", err)
	}

	// prevent zombie
	go func() {
		cmd.Wait()
		globalPHPCmd = nil
	}()

	// verify port is open (wait up to 2 seconds)
	for i := 0; i < 20; i++ {
		if isPortOpen(port) {
			globalPHPCmd = cmd
			globalPHPPort = port
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}

	// if not open → fail
	globalPHPCmd.Process.Kill()
	globalPHPCmd = nil
	return fmt.Errorf("PHP failed to bind to port %d", port)
}

func StopGlobalPHP(res *resolver.Resolver) error {
	if globalPHPCmd != nil {
		err := globalPHPCmd.Process.Kill()
		globalPHPCmd = nil
		return err
	}
	return nil
}

func IsGlobalPHPRunning() bool {
	if globalPHPCmd == nil {
		return false
	}
	return isPortOpen(globalPHPPort)
}

func isPortOpen(port int) bool {
	conn, err := net.Dial("tcp", "127.0.0.1:"+strconv.Itoa(port))
	if err != nil {
		return false
	}
	conn.Close()
	return true
}
