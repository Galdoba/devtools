package cmd

import (
	"fmt"
	"log"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/cls"
	"github.com/Galdoba/devtools/cronex/handler"
	"github.com/urfave/cli/v2"
)

const ()

var logger = cls.New()

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
			logger.AddStdErr(cls.LV_DEBUG, "", 0)
			logger.AddFile(cfg.LogLocation(), cls.LV_DEBUG, "", log.Ldate|log.Ltime)

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
