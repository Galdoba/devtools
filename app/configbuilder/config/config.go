package config

import (
	"fmt"
	"os"
	"path/filepath"

	"encoding/json"
)

const (
	appName = "configbuilder"
)

type configuration struct {
	app	string
	path	string
	isCustom	bool
	Encodings	[]string	`json:"Supported Encodings"` //List of supported encodings for autogenerated config [0 : json][1 : toml][2 : yaml]
	MapData	map[string]float64	`json:"Test Map Data"` // 
	StringData	string	`json:"testString"` // [default : some value]
}

type ConfigFile interface {
	Save()	error
	SaveAs(string)	error
	Load()	(*configuration, error)
	LoadCustom(string)	(*configuration, error)
	Path()	string
	AppName()	string
	IsCustom()	bool
}

type Config interface {
	ConfigFile
	Encodings()	[]string
	MapData()	map[string]float64
	StringData()	string
}


////////////NEW-SAVE-LOAD////////////
//New - autogenerated constructor of config file
func New() *configuration {
	cfg := configuration{}
	cfg.path = `C:\Users\Admin\.config\configbuilder\config.json`
	cfg.app = appName
	cfg.MapData = make(map[string]float64)
	return &cfg
}

//Save - autogenerated constructor of config file
func (cfg *configuration) Save() error {
	data := []byte(header())
	bt, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("can't marshal config: %v", err.Error())
	}
	f, err := os.OpenFile(cfg.path, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return fmt.Errorf("can't open file: %v", err.Error())
	}
	defer f.Close()
	f.Truncate(0)
	data = append(data, bt...)
	_, err = f.Write(data)
	if err != nil {
		return fmt.Errorf("can't write file: %v", err.Error())
	}
	return nil
}

//SaveAs - autogenerated saver of alternative config file
func (cfg *configuration) SaveAs(path string) error {
	cfg.path = path
	cfg.isCustom = true
	return cfg.Save()
}

//Load - Load default config
func Load() (*configuration, error) {
	path := stdConfigPath()
	cfg, err := loadConfig(path)
	if err != nil {
		return nil,  fmt.Errorf("can't load default config: %v", err.Error())
	}
	cfg.isCustom = true
	return cfg, nil
}

//LoadCustom - Loader custom config
func LoadCustom(path string) (*configuration, error) {
	cfg, err := loadConfig(path)
	if err != nil {
		return nil,  fmt.Errorf("can't load custom config: %v", err.Error())
	}
	cfg.isCustom = true
	return cfg, nil
}

//loadConfig - autogenerated loader config file
func loadConfig(path string) (*configuration, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	cfg := &configuration{}
	err = json.Unmarshal(bt, cfg)
	if err != nil {
		return nil, fmt.Errorf("%v", err.Error())
	}
	return cfg, nil
}

////////////HELPERS////////////

//Path - return filepath of current config
func (cfg *configuration) Path() string {
	return cfg.path
}

//IsCustom - return true if config is custom
func (cfg *configuration) IsCustom() bool {
	return cfg.isCustom
}

//AppName - return true if config is custom
func (cfg *configuration) AppName() string {
	return cfg.app
}

func stdConfigDir() string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	path += sep
	return path + ".config" + sep + appName + sep
}

func stdConfigPath() string {
	return stdConfigDir() + "config.json"
}

////////////GETTERS////////////

//GetSupportedEncodings - autogenerated getter for 'Supported Encodings' option of config file
func (cfg *configuration) GetSupportedEncodings()	[]string {
	return cfg.Encodings
}


//GetTestMapData - autogenerated getter for 'Test Map Data' option of config file
func (cfg *configuration) GetTestMapData()	map[string]float64 {
	return cfg.MapData
}


//GetTestString - autogenerated getter for 'testString' option of config file
func (cfg *configuration) GetTestString()	string {
	return cfg.StringData
}

func (cfg *configuration) SetDefault() error {
	cfg.Encodings = append(cfg.Encodings, "json")
	cfg.Encodings = append(cfg.Encodings, "toml")
	cfg.Encodings = append(cfg.Encodings, "yaml")
	cfg.StringData = "some value"
return cfg.Save()}

func header() string {
	hdr := ""
	hdr += `################################################################################` + "\n"
	hdr += `#              This file was generated by 'configbuilder' package              #` + "\n"
	hdr += `#                     Check formatting rules before editing                    #` + "\n"
	hdr += `#      https://www.freecodecamp.org/news/what-is-json-a-json-file-example/     #` + "\n"
	hdr += `################################################################################` + "\n"
	hdr += `# expected location: C:\Users\Admin\.config\configbuilder\config.json` + "\n"
	hdr += `# app name         : configbuilder` + "\n"
	hdr += `` + "\n"
	return hdr
}
