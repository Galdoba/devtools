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

func WorkingDirectoryValid() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory")
	}
	if err := findMainGo(dir); err != nil {
		return err
	}
	return nil
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

func FindVersionControlToken(path string) (int, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return -1, fmt.Errorf("failed to read file: %v")
	}
	lines := strings.Split(string(bt), "\n")
	candidates := []int{}
	for i, line := range lines {
		lowLine := strings.ToLower(line)
		if strings.Contains(lowLine, "version = ") && strings.Contains(lowLine, "//#gvc: version control token") {
			candidates = append(candidates, i)
		}
	}
	switch len(candidates) {
	case 0:
		return -1, fmt.Errorf("'#gvc: version control token' was not found")
	case 1:
		return candidates[0], nil
	default:
		return -1, fmt.Errorf("multiple version control token's was found")
	}
}
