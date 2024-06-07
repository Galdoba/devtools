package cmd

import (
	"fmt"

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
		//Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
		Action: func(c *cli.Context) error {
			_, err := config.Load()
			if err != nil {
				fmt.Printf("config load error: %v\n", err.Error())
				return err
			}
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
}
