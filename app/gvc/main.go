package main

import (
	"os"

	"github.com/Galdoba/devtools/app/gvc/check"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.1.0" //#gvc: version control token
	app.Name = "gvc"
	app.Usage = "wrapper for 'go install' and 'go build' commands"
	app.Description = "gvc is a version control tool for small apps, filling change_log and saving copy of the file for every major and minor upgrade. In order to use all features gvc should me started in a directory with main.go file."
	app.Flags = []cli.Flag{}

	app.Before = func(c *cli.Context) error {
		//check working directory
		return check.WorkingDirectoryValid()
	}
	app.Commands = []*cli.Command{
		//gvc update
		//gvc upgradeMinor
		//gvc upgradeMajor
		//gvc changes
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err := app.Run(args); err != nil {
		println(err.Error())
	}

}
