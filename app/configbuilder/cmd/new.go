package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
	"github.com/Galdoba/devtools/app/configbuilder/internal/tui"
	"github.com/Galdoba/devtools/helpers"
	"github.com/urfave/cli/v2"
)

const (
	action_ADD    = "Add New Field"
	action_EDIT   = "Edit Field"
	action_DELETE = "Delete Field"
	action_DONE   = "DONE"
	sourceName    = "Sourcename"
	dataType      = "DataType"
	designation   = "Designation"
	omit          = "OmitEmpty"
	comment       = "Comment"
	defaultVal    = "Default Value"
)

var modelLanguage string

func NewModel() *cli.Command {
	cmnd := &cli.Command{
		Name:    "new",
		Aliases: []string{},
		Usage:   "create config model",
		Description: strings.Join([]string{
			"interactive loop creates model structure, which will be encoded and saved as model.csv in runtime directory",
			"*this command is expeced to be run in directory with name: '.../config/'",
		}, "\n"),
		Action: func(c *cli.Context) error {
			if modelFileDetected() {
				return fmt.Errorf("can't create new model: %v exist\nsuggestion: run 'configbuilder edit' to change model or 'configbuilder delete' to delete it", configbuilder.MODEL_FILE)
			}
			workingDir, err := filepath.Abs(".")
			if err != nil {
				return err
			}
			if err := configbuilder.CheckWorkingDirectory(workingDir); err != nil {
				if !c.Bool("testmode") {
					return err
				}
			}
			modelLanguage = userSelect("Choose encoding for configfile:", "yaml", "toml", "json")

			cb := configbuilder.New(modelLanguage)
			err = cb.SetSourceDir(workingDir)
			if err != nil {
				return err
			}
			if err := editModel(cb); err != nil {
				return err
			}
			if !c.Bool("testmode") {
				if err = savetoFile(cb.Model()); err != nil {
					return fmt.Errorf("can't save model: %v", err.Error())
				}
			}
			return nil
		},
		Flags: []cli.Flag{},
	}

	return cmnd
}

func modelFileDetected() bool {
	if f, err := os.Stat(configbuilder.MODEL_FILE); err == nil {
		if f.IsDir() {
			return false
		}
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}
	return false
}

func addFieldAction(cb configbuilder.Builder) error {
	f := model.NewField(modelLanguage)
	f.WithSource(userInput("enter Field SourceName:"))
	f.WithDataType(userInput(fmt.Sprintf("SourceName: %v\nenter Field DataType:", f.SourceName)))
	f.WithDesignation(userInput(fmt.Sprintf("SourceName: %v\nDataType: %v\nenter Field Designation:", f.SourceName, f.DataType)))
	return cb.AddField(f)
}

func draw(cb configbuilder.Builder) {
	helpers.ClearTerminal()
	fmt.Println(tui.Status(cb.Model()))
}
