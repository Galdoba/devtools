package cmd

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/urfave/cli/v2"
)

func BuildSource() *cli.Command {
	cmnd := &cli.Command{
		Name:    "build",
		Aliases: []string{},
		Usage:   "build go source file from available model",
		Action: func(c *cli.Context) error {
			if !modelFileDetected() {
				return fmt.Errorf("%v not found\nrun 'configbuilder -h' for help", configbuilder.MODEL_FILE)
			}
			cb, err := loadFromFile()
			if err != nil {
				return fmt.Errorf("can't load model: %v", err.Error())
			}
			err = cb.SetSourceDir(".")
			if err != nil {
				return err
			}
			return buildSources(cb)
		},
		Flags: []cli.Flag{
			// &cli.StringFlag{},
		},
	}
	return cmnd
}

func detectedError(cb configbuilder.Builder) error {
	for _, f := range cb.Model().Fields {
		if err := f.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func buildSources(cb configbuilder.Builder) error {
	sourcetext, err := cb.GenerateSource()
	if err != nil {
		return fmt.Errorf("can't generate source: %v", err.Error())
	}
	f, err := os.OpenFile("config.go", os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("can't open file: %v", err.Error())
	}
	f.Truncate(0)
	if _, err = f.WriteString(sourcetext); err != nil {
		return fmt.Errorf("can't write file: %v", err.Error())
	}
	return nil
}
