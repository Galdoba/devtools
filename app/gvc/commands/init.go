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

		Action: func(c *cli.Context) error {
			_, err := version.Load(WorkingDir + gvc_file)
			if err == nil {
				return fmt.Errorf("gvc file already exists")
			}
			if err := version.New(WorkingDir+gvc_file,
				version.WithName(c.String("project")),
				version.WithDescription(c.String("description")),
			).Save(); err != nil {
				return fmt.Errorf("failed to initiate gvc file: %v", err)
			}
			fmt.Println("gvc file created:", WorkingDir+gvc_file)
			//go to project root root
			//create ./docs
			//create ./docs/gvc_name.md
			//fill basic info
			return nil
		},

		Subcommands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "project",
				Category:    "",
				DefaultText: "",
				FilePath:    "",
				Usage:       "add project name for documentation",
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "description",
				Category:    "",
				DefaultText: "",
				FilePath:    "",
				Usage:       "add project description for documentation",
				Required:    false,
			},
		},
		SkipFlagParsing:        false,
		HideHelp:               false,
		HideHelpCommand:        false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}

}
