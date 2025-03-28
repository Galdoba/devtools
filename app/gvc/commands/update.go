package commands

import (
	"fmt"

	"github.com/Galdoba/devtools/app/gvc/commands/autodoc"
	"github.com/Galdoba/devtools/app/gvc/commands/check"
	"github.com/Galdoba/devtools/app/gvc/commands/inject"
	"github.com/Galdoba/devtools/version"
	"github.com/urfave/cli/v2"
)

func Update() *cli.Command {
	return &cli.Command{
		Name:        "update",
		Aliases:     []string{},
		Usage:       "Create new technical build",
		UsageText:   "gvc update",
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
			found, err := check.GVCfile()
			if err != nil {
				return err
			}
			notes := c.StringSlice("notes")
			switch found {
			case false:
				fmt.Println("No version control file for this project.\nRun 'gvc init' to create one.")
			case true:
				v, err := version.Load(WorkingDir + gvc_file)
				if err != nil {
					return fmt.Errorf("load version failed: %v", err)
				}
				v.Update()
				if err := inject.Inject(v, WorkingDir+main_go_file); err != nil {
					return fmt.Errorf("source injection failed: %v", err)
				}
				if err := v.Save(); err != nil {
					return fmt.Errorf("update failed: %v", err)
				}
				fmt.Printf("update successful\n")
				fmt.Printf("current version: %v\n", v.String())
				amd, err := autodoc.Load(v)
				if err != nil {
					return fmt.Errorf("failed to update docs: %v", err)
				}
				amd.Update(notes...)
				if err = amd.Save(); err != nil {
					return err
				}

			}
			return nil
		},

		Subcommands: []*cli.Command{},
		Flags: []cli.Flag{
			&cli.StringSliceFlag{
				Name:    "notes",
				Usage:   "",
				Aliases: []string{"n"},
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
