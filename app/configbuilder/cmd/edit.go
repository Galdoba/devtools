package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/Galdoba/devtools/app/configbuilder/internal/configbuilder"
	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
	"github.com/urfave/cli/v2"
)

func EditModel() *cli.Command {
	cmnd := &cli.Command{
		Name:    "edit",
		Aliases: []string{},
		Usage:   "edit config model",
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
		Flags: []cli.Flag{
			// &cli.StringFlag{},
		},
	}
	return cmnd
}

func editModel(cb configbuilder.Builder) error {
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

func editFieldAction(cb configbuilder.Builder) error {
	fields := []string{}
	for _, f := range cb.Model().Fields {
		fields = append(fields, f.SourceName)
	}
	toEdit := userSelect("Select field:", fields...)
	f := &model.Field{}
	for _, fCur := range cb.Model().Fields {
		if fCur.SourceName == toEdit {
			f = fCur
		}
	}
	endEdit := false
	for !endEdit {
		switch userSelect(fmt.Sprintf("Editing %v\nWhat must be edited?", f.SourceName), sourceName, dataType, designation, omit, comment, defaultVal, "DONE") {
		case sourceName:
			f.WithSource(userInput("enter Field SourceName:", f.SourceName))
		case dataType:
			f.WithDataType(userInput("enter Field DataType:", f.DataType))
		case designation:
			f.WithDesignation(userInput("enter Field Designation:", f.Designation))
		case omit:
			f.WithOmitEmpty(userConfirm("Omit this Field if empty?"))
		case comment:
			f.WithComment(userInput("enter Field Comment:", f.Comment))
		case defaultVal:
			switch userSelect(fmt.Sprintf("Editing %v\nDefault values actions:", f.SourceName), "add", "edit", "delete") {
			case "add":
				key := "default"
				if userConfirm("Does value have a key?") {
					key = userInput("enter key for value:")
				}
				val := userInput(fmt.Sprintf("enter %v value:", key))
				f.WithValue(key, val)
			case "edit":
				keys := []string{}
				for k := range f.DefaulValDictionary {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				key := userSelect(fmt.Sprintf("Editing %v\nselect key:", f.SourceName), keys...)
				val := userInput(fmt.Sprintf("Editing %v\nenter new value:", f.SourceName), f.DefaulValDictionary[key])
				f.WithValue(key, val)
			case "delete":
				keys := []string{}
				for k := range f.DefaulValDictionary {
					keys = append(keys, k)
				}
				sort.Strings(keys)
				key := userSelect(fmt.Sprintf("Editing %v\nselect key:", f.SourceName), keys...)
				f.DeleteValue(key)

			}
		case "DONE":
			endEdit = true
		}
		draw(cb)
	}
	return nil
}

func deleteFieldAction(cb configbuilder.Builder) error {
	fields := []string{}
	for _, f := range cb.Model().Fields {
		fields = append(fields, f.SourceName)
	}
	toEdit := userSelect("Select field:", fields...)
	index := -1
	for i, fCur := range cb.Model().Fields {
		if fCur.SourceName == toEdit {
			index = i
			break
		}
	}
	return cb.Model().Delete(index)
}

func loadFromFile() (configbuilder.Builder, error) {
	workingDir, err := filepath.Abs(".")
	if err != nil {
		return nil, fmt.Errorf("filepath: %v", err.Error())
	}
	if err := configbuilder.CheckWorkingDirectory(workingDir); err != nil {
		return nil, err
	}
	f, err := os.OpenFile(configbuilder.MODEL_FILE, os.O_RDWR, 0777)
	if err != nil {
		return nil, fmt.Errorf("can't open %v", configbuilder.MODEL_FILE)
	}
	bt, err := os.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("can't read %v: %v", configbuilder.MODEL_FILE, err.Error())
	}
	data := string(bt)

	m, err := model.FromString(data)
	cb := configbuilder.New("any", m)
	return cb, err
}

func savetoFile(m *model.Model) error {
	modelText := m.String()
	f, err := os.OpenFile(configbuilder.MODEL_FILE, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	defer f.Close()
	if err = f.Truncate(0); err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	if _, err = f.WriteString(modelText); err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	return nil
}
