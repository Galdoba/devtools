package cmd

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/urfave/cli/v2"
)

func DeleteModel() *cli.Command {
	cmnd := &cli.Command{
		Name:    "delete",
		Aliases: []string{},
		Usage:   "delete config model",
		Action: func(c *cli.Context) error {
			if !modelFileDetected() {
				fmt.Println(configbuilder.MODEL_FILE, "not found")
				return nil
			}
			if err := os.Remove(configbuilder.MODEL_FILE); err != nil {
				return fmt.Errorf("can't delete model: %v", configbuilder.MODEL_FILE)
			}
			fmt.Println(configbuilder.MODEL_FILE, "deleted")
			return nil
		},
		Flags: []cli.Flag{
			// &cli.StringFlag{},
		},
	}
	return cmnd
}
