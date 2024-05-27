package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
	"github.com/Galdoba/devtools/app/configbuilder/internal/tui"
	"github.com/urfave/cli/v2"
)

func RestoreModel() *cli.Command {
	cmnd := &cli.Command{
		Name:    "restore",
		Aliases: []string{},
		Description: fmt.Sprintf(strings.Join([]string{
			fmt.Sprintf("Read %v in runtime directory and generate %v", configbuilder.SOURCE_FILE, configbuilder.MODEL_FILE),
		}, "\n")),
		Usage:     "Restore model",
		UsageText: "configbuilder restore",
		Action: func(c *cli.Context) error {
			if modelFileDetected() {
				return fmt.Errorf("%v found\nrun 'configbuilder -h' for help", configbuilder.MODEL_FILE)
			}
			cb, err := loadFromSource()
			if err != nil {
				return fmt.Errorf("can't load model: %v", err.Error())
			}
			err = cb.Setup(".", c.App.Version)
			if err != nil {
				return err
			}
			// if err := editModel(cb); err != nil {
			// 	return err
			// }
			if !c.Bool("testmode") {
				if err = savetoFile(cb.Model()); err != nil {
					return fmt.Errorf("can't save model: %v", err.Error())
				}
			}
			return nil
		},
		Flags: []cli.Flag{
			// &cli.StringFlag{},
		},
	}
	return cmnd
}

func restoreModel(cb configbuilder.Builder) error {
	run := true
	for run {
		draw(cb)
		action := userSelect("What is thy action?", action_ADD, action_EDIT, action_DELETE, action_DONE)
		switch action {
		case action_DONE:
			if detectedError(cb) != nil {
				switch userConfirm("This model contains errors.\nDo you really want to stop creation process?") {
				case true:
					run = false
				case false:
					continue
				}
			}
			run = false
		case action_ADD:
			if err := addFieldAction(cb); err != nil {
				userConfirm(fmt.Sprintf("%v\nContinue?", err.Error()))
			}
		case action_EDIT:
			if err := editFieldAction(cb); err != nil {
				return err
			}
		case action_DELETE:
			if err := deleteFieldAction(cb); err != nil {
				userConfirm(fmt.Sprintf("%v\nContinue?", err.Error()))
			}
		default:
			panic("unexpected action: '" + action + "'")
		}
	}
	return nil
}

func loadFromSource() (configbuilder.Builder, error) {
	workingDir, err := filepath.Abs(".")
	if err != nil {
		return nil, fmt.Errorf("filepath: %v", err.Error())
	}
	if err := configbuilder.CheckWorkingDirectory(workingDir); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(configbuilder.SOURCE_FILE, os.O_RDWR, 0777)
	if err != nil {
		return nil, fmt.Errorf("can't open %v", configbuilder.MODEL_FILE)
	}
	bt, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("can't read %v: %v", configbuilder.MODEL_FILE, err.Error())
	}
	data := string(bt)

	m, err := model.FromSource(data)
	cb := configbuilder.New("any", m)
	fmt.Println("")
	fmt.Println(tui.Status(cb.Model()))
	return cb, err
}
