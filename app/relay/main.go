package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/relay/cmd"
	"github.com/urfave/cli/v2"
)

const (
	programName = "relay"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.2.0"
	app.Name = programName
	app.Usage = "Track command queue and relay operations to other apps"
	app.Description = "as a daemon track particular dir and run commands if all is good"
	app.Flags = []cli.Flag{}
	// An action to execute before any subcommands are run, but after the context is ready
	// If a non-nil error is returned, no subcommands are run
	app.Before = func(c *cli.Context) error {

		return nil
	}
	app.Commands = []*cli.Command{
		cmd.Health(),
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err := app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}
