package configmanager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var ErrNoConfig = fmt.Errorf("no candidates detected")
var ErrCantReadDir = fmt.Errorf("can't read default config dir")

// DefaultConfigPath - seeks file with name "config" for application in standard directory.
// It return error if number of candidates found is not 1.
// StdDir is "[USER_HOME]/.config/[APP_NAME]/"
func DefaultConfigPath(app string) (string, error) {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	path += sep
	dir := path + ".config" + sep + app + sep
	de, err := os.ReadDir(dir)
	if err != nil {
		return "", ErrCantReadDir
	}
	potentials := make(map[string]bool)
	for _, e := range de {
		if strings.HasPrefix(e.Name(), "config.") {
			potentials[e.Name()] = true
		}
	}
	if len(potentials) == 0 {
		return "", ErrNoConfig
	}
	if len(potentials) > 1 {
		return "", fmt.Errorf("more than one candidates detected")
	}
	for k := range potentials {
		return k, nil
	}
	return "", fmt.Errorf("no candidates in map (this is unexpected)")
}

// DefaultConfigDir - return default config dir for application
// StdDir is "[USER_HOME]/.config/[APP_NAME]/"
func DefaultConfigDir(app string) string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	path += sep
	return path + ".config" + sep + app + sep
}
