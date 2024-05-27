package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/gpath"
	"github.com/charmbracelet/huh"
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
			issue := ""
			issueList := []string{}
			cfg, err := config.Load()
			fmt.Println(cfg)
			if err != nil {
				issue = fmt.Sprintf("config health: %v", err.Error())
				return fmt.Errorf(issue)
			}
			if cfg.MessageStorageDirectory() == "" {
				issue = fmt.Sprintf("config health: Message Storage directory is not set")
				fmt.Println(issue)
				issueList = append(issueList, issue)
				if c.Bool("fix") {
					dir := gpath.AppStorageDir(c.App.Name)
					inputComponent := huh.NewInput().
						Title("Enter Message Storage Directory:").
						Value(&dir)
					form := huh.NewForm(huh.NewGroup(inputComponent))
					err := form.Run()
					if err != nil {
						return err
					}
					dir = strings.TrimSuffix(dir, " ")
					fmt.Println(dir, "--")
					cfg.SetMessageStorageDirectory(dir)
					if err := cfg.Save(); err != nil {
						fmt.Println(err.Error())
					}
					fmt.Println(cfg.MessageStorageDirectory())
				}
			}
			_, err = os.ReadDir(cfg.MessageStorageDirectory())
			if err != nil {
				issueList = append(issueList, fmt.Sprintf("config health: message storage directory: %v", err.Error()))
			}
			if cfg.RestartCycle() < 1 {
				issueList = append(issueList, fmt.Sprintf("config health: restart cycle expected to be 1 or more seconds"))
			}
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
