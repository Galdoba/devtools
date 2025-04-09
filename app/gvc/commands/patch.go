package commands

import (
	"fmt"

	"github.com/Galdoba/devtools/app/gvc/commands/inject"
	"github.com/Galdoba/devtools/version"
	"github.com/urfave/cli/v2"
)

func Patch() *cli.Command {
	return &cli.Command{
		Name:        "patch",
		Aliases:     []string{},
		Usage:       "Create new patch release",
		UsageText:   "gvc patch -m [message]",
		Description: "Build will be increased. Date will not be written. No copy will be stored or any other changes made.",
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
				return fmt.Errorf("load version failed: %v", err)
			}
			v.Patch(c.String("m"))
			if err := inject.Inject(v, WorkingDir+main_go_file); err != nil {
				return fmt.Errorf("source injection failed: %v", err)
			}
			if err := v.Save(); err != nil {
				return fmt.Errorf("update failed: %v", err)
			}
			fmt.Printf("patch release successful\n")
			fmt.Printf("current version: %v\n", v.String())

			return nil
		},

		Subcommands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:     "message",
				Usage:    "",
				Required: true,
				Aliases:  []string{"m"},
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
