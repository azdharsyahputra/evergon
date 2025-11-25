package pid

import (
	"os"
	"strconv"
)

func Write(path string, pid int) error {
	return os.WriteFile(path, []byte(strconv.Itoa(pid)), 0644)
}

func Read(path string) (int, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	n, err := strconv.Atoi(string(b))
	if err != nil {
		return 0, err
	}
	return n, nil
}

func Remove(path string) {
	os.Remove(path)
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}
