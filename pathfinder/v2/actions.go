package v2

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreatePath - create path assosiated with pathfinder
func (pf *pathfinder) CreatePath() error {
	dir := filepath.Dir(pf.path)
	switch pf.mustExistDir {
	case true:
		if err := validDir(dir); err != nil {
			return err
		}
	case false:
		if err := os.MkdirAll(dir, pf.perm); err != nil {
			return fmt.Errorf("failed to create directory: %v", err)
		}
	}

	if _, err := os.Create(pf.path); err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	return nil
}

// Validate - validate path pathfinder's rules
func (pf *pathfinder) Validate(path string) error {
	return validatePath(*pf, path)
}

// Read - read file assosiated with pathfinder
func (pf *pathfinder) Read() ([]byte, error) {
	if pf.isDir {
		return []byte{}, fmt.Errorf("pathfinder is assosiated with directory: use pf.ReadDir instead")
	}
	return os.ReadFile(pf.path)
}

// ReadDir - read directory pathfinder is assosiated with.
func (pf *pathfinder) ReadDir() ([]os.DirEntry, error) {
	switch pf.isDir {
	case true:
		return os.ReadDir(pf.path)
	case false:
		dir := filepath.Dir(pf.path)
		return os.ReadDir(dir)
	}
	return nil, fmt.Errorf("WUT???")
}

// Write - write file assosiated with pathfinder
func (pf *pathfinder) Write(bt []byte, flag int) error {
	if pf.isDir {
		return fmt.Errorf("pathfinder is assosiated with directory: writing impossible")
	}
	if pf.mustExistFile {
		if err := pathValidation(pf.path); err != nil {
			return fmt.Errorf("writing failed: %v", err)
		}
	}
	f, err := os.OpenFile(pf.path, flag, pf.perm)
	if err != nil {
		return fmt.Errorf("writing failed: %v", err)
	}
	if _, err := f.Write(bt); err != nil {
		return fmt.Errorf("writing failed: %v", err)
	}
	return f.Close()
}
