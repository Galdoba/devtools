package configbuilder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/app/configbuilder/internal/model"
)

const (
	MODEL_FILE = "config_model.csv"
)

func init() {}

// type configBuilder struct {
// 	pathToConfig string
// 	sourcePath   string
// 	app          string
// 	fields       []configField
// }

type configBuilder struct {
	pathToConfig string
	sourcePath   string
	app          string
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
	SetSourceDir(string) error
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

func (cb *configBuilder) SetSourceDir(dir string) error {
	if err := CheckWorkingDirectory(dir); err != nil {
		return err
	}
	dir, _ = filepath.Abs(dir)
	sep := string(filepath.Separator)
	dir = strings.TrimSuffix(dir, sep)
	blocks := strings.Split(dir, sep)
	app := blocks[len(blocks)-2]
	cb.sourcePath = dir + sep + "config.go"
	cb.app = app
	cb.pathToConfig = StdConfigPath(cb.app, cb.language)
	return nil
}

func CheckWorkingDirectory(dir string) error {

	st, err := os.Stat(dir)
	if err != nil {
		return fmt.Errorf("set source error: %v", err.Error())
	}
	if !st.IsDir() {
		return fmt.Errorf("set source error: %v is not dir", dir)
	}
	dir, _ = filepath.Abs(dir)
	sep := string(filepath.Separator)
	dir = strings.TrimSuffix(dir, sep)
	blocks := strings.Split(dir, sep)
	if len(blocks) < 3 {
		return fmt.Errorf("set source error: dir expected to have at least 2 blocks")
	}
	if blocks[len(blocks)-1] != "config" {
		return fmt.Errorf("set source error: dir name must be 'config' (have '%v')", blocks[len(blocks)-1])
	}
	return nil
}

func (cb *configBuilder) AddField(cf *model.Field) error {
	// if err := cf.Validate(); err != nil {
	// 	return fmt.Errorf("can't add field: %v", err.Error())
	// }
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
	// str += `	"path"` + "\n"
	str += `	"path/filepath"` + "\n"
	str += "\n"
	//str += `	"github.com/Galdoba/ffstuff/pkg/configbuilder"` + "\n"
	str += fmt.Sprintf(`	%v`+"\n", packageName)

	str += ")\n"
	str += "\n"
	str += "const (\n"
	str += fmt.Sprintf("	appName = "+`"`+"%v"+`"`+"\n", cb.app)
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
	str += "	Load()	(*configuration, error)\n"
	str += "	LoadCustom(string)	(*configuration, error)\n"
	str += "	Path()	string\n"
	str += "	AppName()	string\n"
	str += "	IsCustom()	bool\n"
	str += "}\n"
	str += "\n"

	str += "type Config interface {\n"
	str += "	ConfigFile\n"
	for _, cf := range cb.model.Fields {
		str += fmt.Sprintf("	%v()	%v\n", cf.SourceName, cf.DataType)
	}
	str += "}\n"
	str += "\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "\n////////////NEW-SAVE-LOAD////////////\n"
	/////////////////////////////////////////////////////////////////////////////
	str += "//New - autogenerated constructor of config file\n"
	str += "func New() *configuration {\n"
	str += "	cfg := configuration{}\n"
	str += fmt.Sprintf("	cfg.path = `%v`\n", cb.pathToConfig)
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
	str += "func Load() (*configuration, error) {\n"
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
	str += "func LoadCustom(path string) (*configuration, error) {\n"
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
		words := strings.Fields(cf.Designation)
		for i := range words {
			letters := strings.Split(words[i], "")
			letters[0] = strings.ToUpper(letters[0])
			words[i] = strings.Join(letters, "")
		}
		funcName := strings.Join(words, "")
		str += fmt.Sprintf("//Get%v - autogenerated getter for '%v' option of config file\n", funcName, strings.TrimSpace(cf.Designation))
		str += fmt.Sprintf("func (cfg *configuration) Get%v()	%v {\n", funcName, cf.DataType)
		str += fmt.Sprintf("	return cfg.%v\n", cf.SourceName)
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

// AutoScan          bool     `yaml:"Scan All Files in Tracked Directories  ,omitempty"`
// type configField struct {
// 	sourceField string
// 	valueType   string
// 	designation string
// 	omitempty   bool
// 	comment     string
// 	dictionary  map[string]string
// }

// func NewField(sourceField, valueType, designation string, omitempty bool) (configField, error) {
// 	if len(sourceField) == 0 {
// 		return configField{}, fmt.Errorf("'sourcefield' must not be blank")
// 	}
// 	switch strings.Split(sourceField, "")[0] {
// 	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
// 	default:
// 		return configField{}, fmt.Errorf("'sourcefield' must start from large latin letter")
// 	}
// 	badCharacters := " `!@#$%^&*()-+|'№;%:?	<>" + `"`
// 	for _, chr := range strings.Split(badCharacters, "") {
// 		if strings.Contains(sourceField, chr) {
// 			return configField{}, fmt.Errorf("can't have character '%v' in 'sourcefield'", chr)
// 		}
// 	}
// 	switch {
// 	default:
// 		return configField{}, fmt.Errorf("unknown primitive '%v' for valueType (TODO: expand valueType to Support interfaces)", valueType)
// 	case TypeValid(valueType):
// 	}

// 	//TODO
// 	return configField{sourceField, valueType, designation, omitempty}, nil
// }

// func (cf *configField) String() string {
// 	oe := ""
// 	if cf.omitempty {
// 		oe = ",omitempty"
// 	}
// 	return fmt.Sprintf("%v	%v	`yaml:"+`"%v%v"`+"`", cf.sourceField, cf.valueType, cf.designation, oe)
// }

//////////HELPERS

// func TypeValid(tp string) bool {
// 	primitives := []string{"bool", "uint8", "uint16", "uint32", "uint64", "int8", "int16", "int32", "int64", "float32", "float64", "complex64", "complex128", "string", "int", "uint", "uintptr", "byte", "rune"}
// 	allTypes := append([]string{}, primitives...)
// 	for _, keys := range primitives {
// 		allTypes = append(allTypes, "[]"+keys)
// 		for _, vals := range primitives {
// 			allTypes = append(allTypes, "map["+keys+"]"+vals)
// 		}
// 	}
// 	for _, t := range allTypes {
// 		if t == tp {
// 			return true
// 		}
// 	}

// 	return false
// }

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
	s += "#              This file was generated by 'configbuilder' package              #\n"
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
	s += `# expected location: ` + cb.pathToConfig + "\n"
	s += fmt.Sprintf("# app name         : %v\n", cb.app)
	return s
}
