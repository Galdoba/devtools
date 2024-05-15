package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/config"
	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
	"github.com/Galdoba/devtools/app/configbuilder/internal/tui"
	"github.com/Galdoba/devtools/helpers"
	"github.com/urfave/cli/v2"
)

const (
	action_ADD          = "Add New Entry"
	action_EDIT         = "Edit Fields"
	action_DELETE       = "Delete Entry"
	action_SWITCH_PLACE = "Switch Entries"
	action_DONE         = "DONE"
	sourceName          = "Sourcename"
	dataType            = "DataType"
	designation         = "Designation"
	omit                = "OmitEmpty"
	comment             = "Comment"
	defaultVal          = "Default Value"
	flag_Testmode       = "testmode"
)

var modelLanguage string

func NewModel() *cli.Command {
	cmnd := &cli.Command{
		Name:      "new",
		Aliases:   []string{},
		Usage:     "Create new config model",
		UsageText: "configbuilder new [command options]",
		Description: strings.Join([]string{
			"Interactive loop creates model structure, which will be encoded and saved as model.csv in runtime directory",
			"*this command is expeced to be run in directory with name: '.../config/'",
		}, "\n"),
		Action: func(c *cli.Context) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("can't create new model: %v", err.Error())
			}
			if modelFileDetected() && !c.Bool(flag_Overwrite) {
				if !c.Bool(flag_Testmode) {
					return fmt.Errorf("%v exist: set flag 'overwrite' to 'true' or run 'configbuilder edit' instead", configbuilder.MODEL_FILE)
				}
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

			modelLanguage = userSelect("Choose encoding for configfile:", cfg.GetSupportedEncodings()...)

			if modelLanguage == "DONE" {
				return fmt.Errorf("can't create new model: encoding was not selected")
			}
			cb := configbuilder.New(modelLanguage)
			err = cb.SetSourceDir(workingDir)
			if err != nil {
				return err
			}
			if err := editModel(cb); err != nil {
				return err
			}
			if !c.Bool(flag_Testmode) {
				if err = savetoFile(cb.Model()); err != nil {
					return fmt.Errorf("can't save model: %v", err.Error())
				}
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "overwrite",
				Usage:   "allow overwrite of existing model's file",
				Aliases: []string{"o"},
			},
		},
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
