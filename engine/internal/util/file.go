package util

import (
	"os"
	"strings"
)

func ReplaceAll(template string, data map[string]string) string {
	for key, val := range data {
		template = strings.ReplaceAll(template, "{{"+key+"}}", val)
	}
	return template
}

func WriteFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
