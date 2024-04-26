package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/urfave/cli/v2"
)

const (
	flag_Overwrite    = "overwrite"
	flag_Delete_model = "delete_model"
)

func BuildSource() *cli.Command {
	cmnd := &cli.Command{
		Name:        "build",
		Aliases:     []string{},
		Usage:       "Generate source file",
		UsageText:   "configbuilder build [options]",
		Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
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
			if sourceFileDetected() && !c.Bool(flag_Overwrite) {
				return fmt.Errorf("can't build source file: flag 'overwrite' is false")
			}
			if err := buildSources(cb); err != nil {
				return fmt.Errorf("can't build source file: %v", err.Error())
			}
			if c.Bool(flag_Delete_model) {
				return removeModel()
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        flag_Overwrite,
				Category:    "",
				DefaultText: "",
				FilePath:    "",
				Usage:       fmt.Sprintf("allow overwrite %v", configbuilder.SOURCE_FILE),
				Aliases:     []string{"o"},
			},
			&cli.BoolFlag{
				Name: flag_Delete_model,

				Usage:   fmt.Sprintf("delete %v after building %v", configbuilder.MODEL_FILE, configbuilder.SOURCE_FILE),
				Aliases: []string{"d"},
			},
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

func sourceFileDetected() bool {
	if f, err := os.Stat(configbuilder.SOURCE_FILE); err == nil {
		if f.IsDir() {
			return false
		}
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}
