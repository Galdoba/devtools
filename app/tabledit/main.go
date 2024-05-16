package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/configbuilder/cmd"
	"github.com/Galdoba/devtools/app/configbuilder/config"
	"github.com/Galdoba/devtools/configmanager"
	"github.com/urfave/cli/v2"
)

const (
	programName = "tabledit"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.0.1"
	app.Name = programName
	app.Usage = "modal csv table edit"
	app.Description = "tabledit have modes: EDIT, SELECT"
	app.Flags = []cli.Flag{}
	// An action to execute before any subcommands are run, but after the context is ready
	// If a non-nil error is returned, no subcommands are run
	app.Before = func(c *cli.Context) error {
		_, err := configmanager.DefaultConfigPath(app.Name)
		if err != nil {
			switch errors.Is(err, configmanager.ErrCantReadDir) {
			case errors.Is(err, configmanager.ErrCantReadDir):
				os.MkdirAll(configmanager.DefaultConfigDir(app.Name), 0777)
			case errors.Is(err, configmanager.ErrNoConfig):
			default:
				return err
			}
			cfg := config.New()
			return cfg.SetDefault()
		}

		return nil
	}
	app.Commands = []*cli.Command{
		cmd.Open(),
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
