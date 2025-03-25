package commands

import (
	"fmt"

	"github.com/Galdoba/devtools/version"
	"github.com/urfave/cli/v2"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Aliases:     []string{},
		Usage:       "Initiate gvc system for current project",
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
			_, err := version.Load(WorkingDir + gvc_file)
			if err == nil {
				return fmt.Errorf("gvc file already exists")
			}
			if err := version.New(WorkingDir + gvc_file).Save(); err != nil {
				return fmt.Errorf("failed to initiate gvc file: %v", err)
			}
			fmt.Println("gvc file created:", WorkingDir+gvc_file)
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
