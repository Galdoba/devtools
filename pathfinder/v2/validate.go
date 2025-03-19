package v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validatePath(pf pathfinder, path string) error {
	dir := filepath.Dir(path)
	if pf.mustExistDir {
		if err := validDir(dir); err != nil {
			return fmt.Errorf("directory confirmation failed: %v", err)
		}
	}
	if err := pathValidation(path); err != nil {
		return fmt.Errorf("path validation failed: %v", err)
	}
	if pf.mustExistFile {
		_, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("path confirmation failed: %v", err)
		}
	}
	if pf.mustHaveExt {
		ext := filepath.Ext(path)
		if "."+pf.ext != ext {
			return fmt.Errorf("path extention expected: %v, but have %v", "."+pf.ext, filepath.Ext(path))
		}
	}
	if pf.mustHaveAppName {
		if !strings.Contains(path, sep+pf.appName+sep) {
			return fmt.Errorf("path expected to have '%v' layer", pf.appName)
		}
	}
	return nil
}
