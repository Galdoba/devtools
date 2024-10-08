package configbuilder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
)

const (
	MODEL_FILE  = "config_model.csv"
	SOURCE_FILE = "config.go"
)

type configBuilder struct {
	pathToConfig string
	sourcePath   string
	app          string
	version      string
	language     string
	model        *model.Model
}

func New(language string, mdl ...*model.Model) *configBuilder {
	cb := configBuilder{}
	cb.language = language
	cb.model = model.NewModel(cb.language)
	for _, m := range mdl {
		cb.model = m
		cb.language = m.Language()
	}
	return &cb
}

type Builder interface {
	Setup(string, string) error
	AddField(*model.Field) error
	DeleteField(int) error
	Model() *model.Model
	GenerateSource() (string, error)
}

func (cb *configBuilder) SetEncoding(s string) error {
	if cb.language != "" {
		return fmt.Errorf("can't set encoding: encoding already set")
	}
	cb.language = s
	return nil
}

//Setup - feed app data to configVulder
//dir = working directory
//version = caller app version
func (cb *configBuilder) Setup(dir, version string) error {
	if err := CheckWorkingDirectory(dir); err != nil {
		return fmt.Errorf("work directory validation: %v", err)
	}
	dir, _ = filepath.Abs(dir)
	sep := string(filepath.Separator)
	dir = strings.TrimSuffix(dir, sep)
	blocks := strings.Split(dir, sep)
	app := blocks[len(blocks)-2]
	cb.sourcePath = dir + sep + SOURCE_FILE
	cb.app = app
	cb.pathToConfig = StdConfigPath(cb.app, cb.language)
	cb.version = version
	return nil
}

func CheckWorkingDirectory(dir string) error {
	st, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("%v", err.Error())
	}
	if !st.IsDir() {
		return fmt.Errorf("%v is not dir", dir)
	}
	dir, _ = filepath.Abs(dir)
	sep := string(filepath.Separator)
	dir = strings.TrimSuffix(dir, sep)
	blocks := strings.Split(dir, sep)
	if len(blocks) < 3 {
		return fmt.Errorf("dir expected to have at least 2 blocks")
	}
	if blocks[len(blocks)-1] != "config" {
		return fmt.Errorf("dir name must be 'config' (have '%v')", blocks[len(blocks)-1])
	}
	return nil
}

func (cb *configBuilder) AddField(cf *model.Field) error {
	for _, flds := range cb.model.Fields {
		if flds.SourceName == cf.SourceName {
			return fmt.Errorf("can't add field: field '%v' is already present", cf.SourceName)
		}
		if flds.Designation == cf.Designation {
			return fmt.Errorf("can't add field: designation '%v' is already present", cf.Designation)
		}
	}
	cb.model.Fields = append(cb.model.Fields, cf)
	return nil
}

func (cb *configBuilder) DeleteField(i int) error {
	return cb.Model().Delete(i)
}

func (cb *configBuilder) Model() *model.Model {
	return cb.model
}

func (cb *configBuilder) GenerateSource() (string, error) {
	encoding := cb.language
	packageName := ""
	marshalFunc := ""
	unmarshalFunc := ""
	encodingExt := ""
	switch encoding {
	case "yaml":
		packageName = `"gopkg.in/yaml.v3"`
		marshalFunc = "yaml.Marshal(cfg)"
		unmarshalFunc = "yaml.Unmarshal(bt, cfg)"
		encodingExt = "yaml"
	case "toml":
		packageName = `"github.com/pelletier/go-toml/v2"`
		marshalFunc = "toml.Marshal(cfg)"
		unmarshalFunc = `toml.Unmarshal(bt, cfg)`
		encodingExt = "toml"
	case "json":
		packageName = `"encoding/json"`
		marshalFunc = `json.MarshalIndent(cfg, "", "  ")`
		unmarshalFunc = `json.Unmarshal(bt, cfg)`
		encodingExt = "json"
	case "":
		return "", fmt.Errorf("encoding is not defined")
	default:
		return "", fmt.Errorf("encoding `%v` is not supported", encoding)
	}
	str := ""
	str += "package config\n"
	str += "\n"
	str += "import (\n"
	str += `	"fmt"` + "\n"
	str += `	"os"` + "\n"
	str += `	"path/filepath"` + "\n"
	str += "\n"
	str += fmt.Sprintf(`	%v`+"\n", packageName)
	str += "\n"
	str += `	"github.com/Galdoba/devtools/configmanager"` + "\n"
	str += ")\n"
	str += "\n"

	str += "const (\n"
	str += fmt.Sprintf("	appName = "+`"`+"%v"+`"`+"\n", cb.app)
	str += fmt.Sprintf("	version = "+`"`+"%v"+`"`+"\n", cb.version)
	str += ")\n"
	str += "\n"

	str += "type configuration struct {\n"
	str += "	app	string\n"
	str += "	path	string\n"
	str += "	isCustom	bool\n"
	for _, cf := range cb.model.Fields {
		str += fmt.Sprintf("	%v\n", cf.String())
	}
	str += "}\n"
	str += "\n"

	str += "type ConfigFile interface {\n"
	str += "	Save()	error\n"
	str += "	SaveAs(string)	error\n"
	str += "	SetDefault()	error\n"
	//str += "	LoadCustom(string)	(Config, error)\n"
	str += "	Path()	string\n"
	str += "	AppName()	string\n"
	str += "	IsCustom()	bool\n"
	str += "}\n"
	str += "\n"

	str += "type Config interface {\n"
	str += "	ConfigFile\n"
	for _, cf := range cb.model.Fields {
		funcName := funcName(cf.Designation)
		str += fmt.Sprintf("	%v()	%v\n", funcName, cf.DataType)
	}
	for _, cf := range cb.model.Fields {
		funcName := funcName(cf.Designation)
		str += fmt.Sprintf("	Set%v(%v)\n", funcName, cf.DataType)
	}
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n////////////NEW-SAVE-LOAD////////////\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//New - autogenerated constructor of config file\n"
	str += "func New() Config {\n"
	str += "	cfg := configuration{}\n"
	// str += fmt.Sprintf("	cfg.path = `%v`\n", cb.pathToConfig)
	str += fmt.Sprintf("	cfg.path = configmanager.DefaultConfigDir(appName) + "+`"config.`+"%v"+`"`+"\n", encodingExt)
	// func DefaultConfigDir(app string) string {
	// 	path, err := os.UserHomeDir()
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}
	// 	sep := string(filepath.Separator)
	// 	path += sep
	// 	return path + ".config" + sep + app + sep
	// }

	str += "	cfg.app = appName\n"
	for _, fld := range cb.model.Fields {
		if strings.Contains(fld.DataType, "map[") {
			str += fmt.Sprintf("	cfg.%v = make(%v)\n", fld.SourceName, fld.DataType)
		}
	}
	str += "	return &cfg\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//Save - autogenerated constructor of config file\n"
	str += "func (cfg *configuration) Save() error {\n"
	str += "	data := []byte(header())\n"
	str += fmt.Sprintf("	bt, err := %v\n", marshalFunc)
	str += "	if err != nil {\n"
	str += "		return fmt.Errorf(" + `"` + "can't marshal config: " + `%` + `v` + `"` + ", err.Error())\n"
	str += "	}\n"
	// str += "	dir :=  path.Dir(cfg.path)\n"
	str += "	if err := os.MkdirAll(filepath.Dir(cfg.path), 0777); err != nil {\n"
	str += `		return fmt.Errorf("can't create directory")` + "\n"
	str += "	}\n"
	str += "	f, err := os.OpenFile(cfg.path, os.O_CREATE|os.O_WRONLY, 0777)\n"
	str += "	if err != nil {\n"
	str += "		return fmt.Errorf(" + `"` + "can't open file: " + `%` + `v` + `"` + ", err.Error())\n"
	str += "	}\n"
	str += "	defer f.Close()\n"
	str += "	f.Truncate(0)\n"
	// str += "	_, err = f.WriteString(header())\n"
	// str += "	if err != nil {\n"
	// str += "		return fmt.Errorf(" + `"` + "can't save file: write header: " + `%` + `v` + `"` + ", err.Error())\n"
	// str += "	}\n"
	str += "	data = append(data, bt...)\n"
	str += "	_, err = f.Write(data)\n"
	str += "	if err != nil {\n"
	str += "		return fmt.Errorf(" + `"` + "can't write file: " + `%` + `v` + `"` + ", err.Error())\n"
	str += "	}\n"
	str += "	return nil\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//SaveAs - autogenerated saver of alternative config file\n"
	str += "func (cfg *configuration) SaveAs(path string) error {\n"
	str += "	cfg.path = path\n"
	str += "	cfg.isCustom = true\n"
	str += "	return cfg.Save()\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//Load - Load default config\n"
	str += "func Load() (Config, error) {\n"
	str += "	path := stdConfigPath()\n"
	str += "	cfg, err := loadConfig(path)\n"
	str += "	if err != nil {\n"
	str += "		return nil,  fmt.Errorf(" + `"can't load default config: ` + `%` + `v"` + ", err.Error())\n"
	str += "	}\n"
	str += "	cfg.isCustom = true\n"
	str += "	return cfg, nil\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//LoadCustom - Loader custom config\n"
	str += "func LoadCustom(path string) (Config, error) {\n"
	str += "	cfg, err := loadConfig(path)\n"
	str += "	if err != nil {\n"
	str += "		return nil,  fmt.Errorf(" + `"can't load custom config: ` + `%` + `v"` + ", err.Error())\n"
	str += "	}\n"
	str += "	cfg.isCustom = true\n"
	str += "	return cfg, nil\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//loadConfig - autogenerated loader config file\n"
	str += "func loadConfig(path string) (*configuration, error) {\n"
	str += "	bt, err := os.ReadFile(path)\n"
	str += "	if err != nil {\n"
	str += "		return nil, fmt.Errorf(" + `"%` + `v"` + ", err.Error())\n"
	str += "	}\n"
	str += "	cfg := &configuration{}\n"
	str += "	cfg.path = path\n"
	str += "	cfg.app = appName\n"

	str += fmt.Sprintf("	err = %v\n", unmarshalFunc)
	str += "	if err != nil {\n"
	str += "		return nil, fmt.Errorf(" + `"%` + `v"` + ", err.Error())\n"
	str += "	}\n"
	str += "	return cfg, nil\n"
	str += "}\n"

	/////////////////////////////////////////////////////////////////////////////
	str += "\n////////////HELPERS////////////\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n"
	str += "//Path - return filepath of current config\n"
	str += "func (cfg *configuration) Path() string {\n"
	str += "	return cfg.path\n"
	str += "}\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n"
	str += "//IsCustom - return true if config is custom\n"
	str += "func (cfg *configuration) IsCustom() bool {\n"
	str += "	return cfg.isCustom\n"
	str += "}\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n"
	str += "//AppName - return true if config is custom\n"
	str += "func (cfg *configuration) AppName() string {\n"
	str += "	return cfg.app\n"
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n"
	str += "//DefaultPath - return default path of current config\n"
	str += "func DefaultPath() string {\n"
	str += fmt.Sprintf("	return configmanager.DefaultConfigDir(appName) + "+`"config.`+"%v"+`"`+"\n", encodingExt)
	str += "}\n"

	str += "func stdConfigDir() string {\n"
	str += "	path, err := os.UserHomeDir()\n"
	str += "	if err != nil {\n"
	str += "		panic(err.Error())\n"
	str += "	}\n"
	str += "	sep := string(filepath.Separator)\n"
	str += "	path += sep\n"
	str += "	return path + " + `".config"` + " + sep + appName + sep\n"
	str += "}\n"
	str += "\n"
	str += "func stdConfigPath() string {\n"
	str += "	return stdConfigDir() + " + `"config.` + encodingExt + `"` + "\n"
	str += "}\n"

	/////////////////////////////////////////////////////////////////////////////
	str += "\n////////////GETTERS////////////\n"
	for _, cf := range cb.model.Fields {
		str += "\n"
		funcName := funcName(cf.Designation)
		str += fmt.Sprintf("//%v - autogenerated getter for '%v' option\n", funcName, strings.TrimSpace(cf.Designation))
		for _, line := range wrapText(cf.Comment, 1, 78) {
			str += fmt.Sprintf("//%v\n", line)
		}
		str += fmt.Sprintf("func (cfg *configuration) %v()	%v {\n", funcName, cf.DataType)
		str += fmt.Sprintf("	return cfg.%v\n", cf.SourceName)
		str += "}\n"
		str += "\n"
	}

	str += "\n////////////SETTERS////////////\n"
	for _, cf := range cb.model.Fields {
		str += "\n"
		funcName := funcName(cf.Designation)
		valName := strings.ToLower(cf.SourceName)
		str += fmt.Sprintf("//Set%v - autogenerated setter for '%v' option\n", funcName, strings.TrimSpace(cf.Designation))
		str += fmt.Sprintf("func (cfg *configuration) Set%v(%v %v)	{\n", funcName, valName, cf.DataType)
		str += fmt.Sprintf("	cfg.%v = %v\n", cf.SourceName, valName)
		str += "}\n"
		str += "\n"
	}
	////////////////////////////////////////////////////////////////////////////////

	str += generateDefaultFunc(cb)

	str += "func header() string {\n"
	str += `	hdr := ""` + "\n"
	for _, s := range strings.Split(cb.headerConstruct(), "\n") {
		str += fmt.Sprintf("	hdr += `"+"%v"+"` + "+`"\n"`+"\n", s)
	}
	str += "	return hdr\n"
	str += "}\n"

	return str, nil
}

func generateDefaultFunc(cb Builder) string {
	s := ""
	s += "func (cfg *configuration) SetDefault() error {\n"
field:
	for _, f := range cb.Model().Fields {
		for k, v := range f.DefaulValDictionary {
			if strings.TrimSpace(v) == strings.TrimSpace(k) && strings.TrimSpace(v) == "" {
				continue
			}
			comp, kType, vType := model.DataTypeSegments(f.DataType)
			if kType == model.DateType_string {
				k = `"` + k + `"`
			}
			if vType == model.DateType_string {
				v = `"` + v + `"`
			}
			switch comp {
			case model.DataComposition_MAP:
				s += "	cfg." + f.SourceName + "[" + k + "]" + " = " + v + "\n"
			case model.DataComposition_SLICE:
				s += "	cfg." + f.SourceName + " = append(cfg." + f.SourceName + ", " + v + ")\n"
			case model.DataComposition_PRIMITIVE:
				s += "	cfg." + f.SourceName + " = " + v + "\n"
				continue field
			}
		}
	}
	s += "return cfg.Save()}\n"
	s += "\n"
	return s
}

func StdConfigPath(app, encoding string) string {
	return StdConfigDir(app) + "config." + encoding
}

func StdConfigDir(name string) string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	path += sep
	return path + ".config" + sep + name + sep
}

func (cb *configBuilder) headerConstruct() string {
	s := "################################################################################\n"
	s += "#                 This file was generated by configbuilder app                 #\n"
	s += "#                     Check formatting rules before editing                    #\n"
	switch cb.language {
	case "json":

		s += `#      https://www.freecodecamp.org/news/what-is-json-a-json-file-example/     #` + "\n"
	case "toml":
		s += `#                 https://docs.fileformat.com/programming/toml/                #` + "\n"
	case "yaml":
		s += `#                 https://docs.fileformat.com/programming/yaml/                #` + "\n"
	}
	s += "################################################################################\n"
	s += fmt.Sprintf("# builder version  : %v\n", cb.version)
	s += `# expected location: ` + cb.pathToConfig + "\n"
	s += fmt.Sprintf("# app name         : %v\n", cb.app)

	return s
}

func funcName(designation string) string {
	words := strings.Fields(designation)
	for i := range words {
		letters := strings.Split(words[i], "")
		letters[0] = strings.ToUpper(letters[0])
		words[i] = strings.Join(letters, "")
	}
	funcName := strings.Join(words, "")
	return funcName
}

func wrapText(text string, min, max int) []string {
	lines := []string{}
	if text == "" {
		return lines
	}
	if min < 1 {
		min = 1
	}
	words := strings.Fields(text)
	currentLine := ""
	for _, word := range words {
		if len(currentLine)+len(word)+2 < max {
			currentLine += word + " "
			continue
		}
		switch currentLine {
		case "":
			lines = append(lines, word)
			currentLine = ""
		default:
			lines = append(lines, currentLine)
			currentLine = word + " "
		}
	}
	if currentLine != "" {
		lines = append(lines, currentLine)
	}
	return lines
}
