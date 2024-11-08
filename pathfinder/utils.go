package pathfinder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func homeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	return home + sep

}

var sep = string(filepath.Separator)

func validDir(dir string) error {
	fi, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("failed to open directory: %v", err)
	}
	if !fi.IsDir() {
		return fmt.Errorf("%v is not dir", dir)
	}
	return nil
}

func pathValidation(path string) error {
	semicolonCount := 0
	for _, glyph := range strings.Split(path, "") {
		switch glyph {
		case `<`, `>`, `"`, `|`, `?`, `*`, "\t", "\n", "\r":
			return fmt.Errorf("path contains forbidden character: %v", glyph)
		case ":":
			semicolonCount++
			if semicolonCount != 1 {
				return fmt.Errorf("path contains forbidden character: %v", glyph)
			}
			disk := false
			lowPath := strings.ToLower(path)
			for _, l := range strings.Split("abdcefghijklmnopqrstuvwxyz", "") {
				if strings.HasPrefix(lowPath, l+":") {
					disk = true
					break
				}
			}
			if !disk {
				return fmt.Errorf("path contains forbidden character: %v", glyph)
			}
		default:
		}
	}
	return nil
}
