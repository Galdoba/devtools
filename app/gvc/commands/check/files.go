package check

import (
	"fmt"
	"os"
	"strings"
)

const (
	gvc_file     = "version.gvc"
	main_go_file = "main.go"
)

func GVCfile() (bool, error) {
	dir, err := os.Getwd()
	if err != nil {
		return false, fmt.Errorf("failed to get working directory")
	}
	if err := confirmFile(dir, gvc_file); err != nil {
		if strings.Contains(err.Error(), "was not found") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func MainGoFile() error {
	dir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory")
	}
	if err := confirmFile(dir, main_go_file); err != nil {
		return err
	}
	return nil
}
