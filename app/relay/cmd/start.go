package cmd

import (
	"fmt"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/cronex/handler"
	"github.com/urfave/cli/v2"
)

const ()

func Start() *cli.Command {
	cmnd := &cli.Command{
		Name:      "start",
		Aliases:   []string{},
		Usage:     "Create handler and do Jobs",
		UsageText: "create handler in storage directory which will track job fileles and execute them if time is valid",
		//Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
		Action: func(c *cli.Context) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			handler.Start(cfg.MessageStorageDirectory(),
				handler.Cycle(cfg.RestartCycle()),
			)

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "fix",
				Usage: "try to fix encountered issues",
			},
		},
	}
	return cmnd
}
