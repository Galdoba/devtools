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

/*
run


*/

var configPath string

const (
	programName = "configbuilder"
)

func main() {
	app := cli.NewApp()
	app.Version = "v 0.1.0"
	app.Name = programName
	app.Usage = "fast build config file for go applications"
	// app.UsageText = "configbuilder command"
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:    "testmode",
			Usage:   "DO NOT check directory and DO NOT save model to file",
			Aliases: []string{"t"},
		},
	}
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
		cmd.NewModel(),
		cmd.EditModel(),
		cmd.DeleteModel(),
		cmd.BuildSource(),
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	fmt.Println("Main:")
	if err := app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}
