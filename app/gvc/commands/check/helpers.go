package check

import (
	"fmt"
	"os"
)

func confirmFile(dir, file string) error {
	if dir == "" {
		return fmt.Errorf("can't confirm: directory not provided")
	}
	if file == "" {
		return fmt.Errorf("can't confirm: file not provided")
	}
	fi, err := os.ReadDir(dir)
	if err != nil {
		return fmt.Errorf("failed to read directory")
	}
	for _, f := range fi {
		if f.Name() == file {
			return nil
		}
	}
	return fmt.Errorf("file %v was not found in %v", file, dir)
}
