package commands

import (
	"path/filepath"

	"github.com/Galdoba/devtools/app/gvc/commands/check"
)

var WorkingDir string
var gvc_file = "version.gvc"
var main_go_file = "main.go"
var app_name = ""

func CheckWorkingDirectory() error {
	wd, err := check.WorkingDirectory()
	if err != nil {
		return err
	}
	WorkingDir = wd
	app_name = filepath.Base(filepath.Dir(wd))
	return nil
}
