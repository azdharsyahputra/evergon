package manager

import (
	"fmt"
	"math/rand"
	"time"
)

func GeneratePort() string {
	rand.Seed(time.Now().UnixNano())
	port := 9100 + rand.Intn(400)
	return fmt.Sprintf("%d", port)
}
