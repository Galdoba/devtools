package commands

import "github.com/Galdoba/devtools/app/gvc/commands/check"

var WorkingDir string
var gvc_file = "version.gvc"
var main_go_file = "main.go"

func CheckWorkingDirectory() error {
	wd, err := check.WorkingDirectory()
	if err != nil {
		return err
	}
	WorkingDir = wd
	return nil
}
