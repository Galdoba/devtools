package commands

import (
	"fmt"

	"github.com/Galdoba/devtools/version"
	"github.com/urfave/cli/v2"
)

func Status() *cli.Command {
	return &cli.Command{
		Name:        "status",
		Aliases:     []string{},
		Usage:       "Print current project's gvc file",
		UsageText:   "",
		Description: "",
		Args:        false,
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(*cli.Context) {
		},
		Before: func(*cli.Context) error {
			return CheckWorkingDirectory()
		},

		Action: func(c *cli.Context) error {
			v, err := version.Load(app_name)
			if err != nil {
				return fmt.Errorf("gvc not exists")
			}
			fmt.Printf("%v source code has version %v", app_name, v.String())

			return nil
		},

		Subcommands:            []*cli.Command{},
		Flags:                  []cli.Flag{},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}
