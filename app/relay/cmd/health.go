package cmd

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/urfave/cli/v2"
)

const ()

func Health() *cli.Command {
	cmnd := &cli.Command{
		Name:      "health",
		Aliases:   []string{},
		Usage:     "Check relay's program data",
		UsageText: "relay health",
		Before: func(c *cli.Context) error {
			//health не пишет в лог ибо хз чё там с конфигомы
			return nil
		},
		//Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
		Action: func(c *cli.Context) error {
			cfg, err := config.Load()
			if err != nil {
				fmt.Printf("config load error: %v\n", err.Error())
				return err
			}
			f, err := os.Stat(cfg.MessageStorageDirectory())
			if err != nil {
				fmt.Printf("health: Message Storage Directory: %v", err.Error())
				return err
			}
			if !f.IsDir() {
				fmt.Printf("health: %v is not a directory", err.Error())
				return err
			}
			lf, err := os.Open(cfg.LogLocation())
			if err != nil {
				fmt.Printf("health: log file: %v\n", err.Error())
				return err
			}
			defer lf.Close()
			fmt.Println("health is good")
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
	///test
}
