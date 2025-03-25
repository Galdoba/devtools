package check

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	sep = string(filepath.Separator)
)

func WorkingDirectory() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory")
	}
	if err := findMainGo(dir); err != nil {
		return "", err
	}
	dir = strings.TrimSuffix(dir, sep) + sep
	return dir, nil
}

func findMainGo(dir string) error {
	fi, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read dirirectory")
	}
	for _, f := range fi {
		if f.Name() == "main.go" {
			return nil
		}
	}
	return fmt.Errorf("file main.go was not found in %v", dir)
}
