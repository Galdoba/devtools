package cmd

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/urfave/cli/v2"
)

func DeleteModel() *cli.Command {
	cmnd := &cli.Command{
		Name:        "delete",
		Aliases:     []string{},
		Usage:       "Delete model",
		UsageText:   "configbuilder delete [options]",
		Description: fmt.Sprintf("delete %v in runtime directory", configbuilder.MODEL_FILE),
		Args:        false,
		Action: func(c *cli.Context) error {
			if !modelFileDetected() {
				fmt.Println(configbuilder.MODEL_FILE, "not found")
				return nil
			}
			removeModel()
			fmt.Println(configbuilder.MODEL_FILE, "deleted")
			return nil
		},
		Flags: []cli.Flag{
			// &cli.StringFlag{},
		},
	}
	return cmnd
}

func removeModel() error {
	if err := os.Remove(configbuilder.MODEL_FILE); err != nil {
		return fmt.Errorf("can't delete model: %v", configbuilder.MODEL_FILE)
	}
	return nil
}
