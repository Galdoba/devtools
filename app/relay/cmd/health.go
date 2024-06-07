package cmd

import (
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
