package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/configmanager"
	"github.com/Galdoba/devtools/decidion/operator"
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
			issuesMet := true
			for issuesMet {
				_, err := configmanager.DefaultConfigPath(c.App.Name)
				if err != nil {
					fmt.Println("config health: ", err.Error())
					if !c.Bool("fix") {
						return fmt.Errorf("config confirmation failed")
					}
					switch err {
					case configmanager.ErrCantReadDir:
						cfgDir := configmanager.DefaultConfigDir(c.App.Name)
						switch operator.Confirm(fmt.Sprintf("Create directory '%v'?", cfgDir)) {
						case true:
							if err := os.MkdirAll(cfgDir, 0777); err != nil {
								return fmt.Errorf("can't create config directory: %v", err.Error())
							}
							continue
						case false:
							return fmt.Errorf("config confirmation failed")
						}
					case configmanager.ErrNoConfig:
						switch operator.Confirm(fmt.Sprintf("Create default config?")) {
						case true:
							cfg := config.New()
							cfg.SetDefault()
							if err := cfg.Save(); err != nil {
								return fmt.Errorf("can't save config: %v", err.Error())
							}
							if cfg.MessageStorageDirectory() == "" && operator.Confirm("Set new message storage directory?") {
								dir, err := operator.Input("Enter Message Storage Directory", dirConfirm)
								if err != nil {
									return fmt.Errorf("can't set value")
								}
								cfg.SetMessageStorageDirectory(dir)
							}
						case false:
							return fmt.Errorf("config confirmation failed")
						}

					}
				}

				issuesMet = false
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

func dirConfirm(dir string) error {
	_, err := os.ReadDir(strings.TrimSuffix(dir, " "))
	return err
}
