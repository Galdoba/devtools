package v2

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type pathfinder struct {
	root              string //default: UserHome
	appName           string
	template          string //program/config/log
	systemlayers      []string
	prefixlayers      []string
	suffixlayers      []string
	fileName          string
	ext               string
	perm              os.FileMode
	isDir             bool
	path              string
	mustHaveAppName   bool
	mustHaveSignature bool
	mustHaveExt       bool
	mustExistDir      bool
	mustExistFile     bool
	skipValidation    bool
}

var sep = string(filepath.Separator)

const (
	template_UNDEFINED = "undefined"
	TEMPLATE_PROGRAM   = "program"
	TEMPLATE_DOCUMENT  = "document"
	TEMPLATE_CONFIG    = "config"
	TEMPLATE_LOG       = "log"
	TEMPLATE_CUSTOM    = "custom"
	rootHOME           = "rootHOME"
)

func New(template string, pfOptions ...Option) (*pathfinder, error) {
	pf := pathfinder{}
	pf.root = rootHOME
	pf.template = template
	pf.perm = 0666
	pf = applyTemplate(pf)
	for _, enrich := range pfOptions {
		enrich(&pf)
	}
	pf.path = constructPath(pf)
	if pf.skipValidation {
		return &pf, nil
	}
	if err := validatePath(pf, pf.path); err != nil {
		return nil, err
	}
	return &pf, nil
}

func applyTemplate(pf pathfinder) pathfinder {
	switch pf.template {
	case TEMPLATE_PROGRAM:
		pf.mustHaveSignature = true
		pf.mustHaveAppName = true
		pf.systemlayers = []string{"Programs"}
	case TEMPLATE_LOG:
		pf.mustHaveSignature = true
		pf.mustHaveAppName = true
		pf.ext = "log"
		pf.systemlayers = []string{".log"}
	case TEMPLATE_CONFIG:
		pf.mustHaveSignature = true
		pf.mustHaveAppName = true
		pf.ext = "config"
		pf.systemlayers = []string{".config"}
	case TEMPLATE_DOCUMENT:
		pf.systemlayers = []string{"Documents"}
	case TEMPLATE_CUSTOM:
	default:
		pf.template = template_UNDEFINED
	}
	return pf
}

func constructPath(pf pathfinder) string {
	path := pf.root
	switch pf.root {
	case rootHOME:
		path = homeDir()
	default:
	}
	for _, tech := range pf.systemlayers {
		path += sep + tech
	}
	if pf.mustHaveSignature {
		path += sep + "galdoba"
	}
	for _, pref := range pf.prefixlayers {
		path += sep + pref
	}
	if pf.appName != "" {
		path += sep + pf.appName
	}
	for _, suff := range pf.suffixlayers {
		path += sep + suff
	}
	if pf.fileName != "" {
		path += sep
		path += pf.fileName
		if pf.ext != "" {
			path += "." + pf.ext
		}
	}
	if pf.isDir {
		path += sep
	}

	return path
}

type Option func(*pathfinder)

//root+prefix+appName+suffix+filename+ext

/*
app := "my_app"
pathfinder.NewPath(pathfinder.IsConfig(),pathfinder.WithProgram(app), pathfinder.WithFileName(app+".config"),)

*/

func homeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	return home

}

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
