package main

import (
	"os"

	"github.com/Galdoba/devtools/app/gvc/commands"
	"github.com/Galdoba/devtools/text"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Version = "0.0.1:20250409 [build 6]" //#gvc: version control token
	app.Name = "gvc"
	app.Usage = "wrapper for 'go install' and 'go build' commands"
	app.Description = text.Wrap("gvc is a version control tool for small apps, filling change_log and saving copy of the file for every major and minor upgrade.\nIn order to use all features gvc should me started in a directory with main.go file.",
		text.MaxWidth(77), text.LeftOffset(0), text.WrapLimit(10),
	)
	app.Flags = []cli.Flag{}

	app.Commands = []*cli.Command{
		commands.Init(),
		commands.Status(),
		commands.Major(),
		commands.Minor(),
		commands.Patch(),
		commands.Update(),

		//gvc init
		//gvc health
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
