package commands

import (
	"fmt"

	"github.com/Galdoba/devtools/app/gvc/commands/check"
	"github.com/Galdoba/devtools/version"
	"github.com/urfave/cli/v2"
)

func Stat() *cli.Command {
	return &cli.Command{
		Name:        "stat",
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

		Action: func(*cli.Context) error {
			found, err := check.GVCfile()
			if err != nil {
				return err
			}
			switch found {
			case false:
				fmt.Println("No version control file for this project.\nRun 'gvc init' to create one.")
			case true:
				v, err := version.Load(WorkingDir + gvc_file)
				if err != nil {
					return fmt.Errorf("load version failed: %v", err)
				}
				fmt.Printf("current version: %v\n", v.String())

			}
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
