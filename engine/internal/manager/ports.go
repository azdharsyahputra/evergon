package manager

import (
	"fmt"
	"net"
)

func IsPortAvailable(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}
	ln.Close()
	return true
}

func FindAvailablePort() string {
	for p := 9000; p <= 9999; p++ {
		if IsPortAvailable(fmt.Sprintf("%d", p)) {
			return fmt.Sprintf("%d", p)
		}
	}
	return "0"
}
