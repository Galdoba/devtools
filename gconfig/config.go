package gconfig

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/gpath"
	"gopkg.in/yaml.v3"
)

const (
	TYPE_BOOL   = "BOOL"
	TYPE_INT    = "INT"
	TYPE_FLOAT  = "FLOAT"
	TYPE_STRING = "STRING"
	IS_CONCRETE = "CONCRETE"
	IS_SLICE    = "SLICE"
	IS_MAP      = "MAP"
)

// type Config interface {
// 	Path() string
// 	Data() []byte
// }

/*
#######################################################
#  This is auto generated config file.                #
#  Check formatting rules before manual edit.         #
#  https://docs.fileformat.com/programming/yaml/      #
#######################################################

program: mfline
String Type Parameters:
	default_scan_storage_directory: /home/galdoba/.ffstuff/data/mfline/
	log_file(unimplemented): /home/galdoba/.ffstuff/logs/mfline.log
Float Type Parameters:

*/

// config struct  
type Config struct {
	program           string                       //`yaml:"Application"`
	Tags              []string                     `yaml:"Config tags,omitempty"`
	location          string                       //`yaml:"Config location,omitempty"`
	header            string                       `yaml:"#,omitempty"`
	Option_STR        map[string]string            `yaml:"String Options,omitempty"`
	Option_INT        map[string]int               `yaml:"Integer Options,omitempty"`
	Option_FLOAT64    map[string]float64           `yaml:"Float Options,omitempty"`
	Option_BOOL       map[string]bool              `yaml:"Boolean Options,omitempty"`
	Option_LIST       map[string][]string          `yaml:"Lists,omitempty"`
	Option_DICTIONARY map[string]map[string]string `yaml:"Dictionaries,omitempty"`
}

func (cfg *Config) Location() string {
	return cfg.location
}

func NewConfig(program string, tags ...instruction) (*Config, error) {
	if len(tags) == 0 {
		return nil, fmt.Errorf("can't create config: at least one instruction must be set")
	}

	cfg := Config{}
	cfg.program = program
	cfg.Tags = append(cfg.Tags, "app: "+program)
	loc := stdPath(program)
	for _, t := range tags {
		switch t.operation {
		case toFile:
			loc = t.value
			//cfg.Tags = append(cfg.Tags, "location: "+loc)
		case defaultCase:
			if len(tags) != 1 {
				return nil, fmt.Errorf("can't create config: instruction 'Default' must be only instruction to exist")
			}
			cfg.header = header(cfg.program)
		}

	}
	cfg.location = loc
	cfg.Tags = append(cfg.Tags, "location: "+loc)
	cfg.Option_BOOL = make(map[string]bool)
	cfg.Option_INT = make(map[string]int)
	cfg.Option_FLOAT64 = make(map[string]float64)
	cfg.Option_STR = make(map[string]string)
	cfg.Option_LIST = make(map[string][]string)
	cfg.Option_DICTIONARY = make(map[string]map[string]string, 0)

	return &cfg, nil
}

func (cfg *Config) Save(instructions ...instruction) error {
	location := cfg.location
	for _, inst := range instructions {
		switch inst.operation {
		case toFile:
			location = inst.value
			cfg.Tags = append(cfg.Tags, "location: "+location)
		}
	}
	data := []byte(cfg.header)
	bt, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	dir := filepath.Dir(location)
	if dir == "." {
		return fmt.Errorf("can't define directory for %v", location)
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(location, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer f.Close()
	f.Truncate(0)
	if cfg.header != "" {
		_, err := f.WriteString(cfg.header)
		if err != nil {
			return fmt.Errorf("can't save file: write header: %v", err.Error())
		}
	}

	data = append(data, bt...)
	_, err = f.Write(bt)
	return err
}

func Load(program string, instructions ...instruction) (*Config, error) {
	path := stdPath(program)
	for _, inst := range instructions {
		switch inst.operation {
		case fromFile:
			path = inst.value
		}
	}
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("can't load config: %v", err.Error())
	}
	cfg, err := NewConfig(program, ToFile(path))
	if err != nil {
		return nil, fmt.Errorf("can't load config: %v", err.Error())
	}
	err = yaml.Unmarshal(bt, cfg)
	if err != nil {
		return nil, fmt.Errorf("can't load config: %v", err.Error())
	}
	return cfg, nil
}

func stdPath(program string) string {
	path := gpath.StdConfigDir(program) + "config.yaml"
	return path
}

type instruction struct {
	operation uint
	value     string
}

const (
	fromFile    uint = 100
	toFile      uint = 101
	defaultCase uint = 102
)

func FromFile(path string) instruction {
	return instruction{
		operation: fromFile,
		value:     path,
	}
}

func ToFile(path string) instruction {
	return instruction{
		operation: toFile,
		value:     path,
	}
}

func Default() instruction {
	return instruction{
		operation: defaultCase,
	}
}

func header(program string) string {
	return strings.Join([]string{
		fmt.Sprintf("#######################################################"),
		fmt.Sprintf("#  This is auto generated config file.                #"),
		fmt.Sprintf("#  Check formatting rules before manual edit.         #"),
		fmt.Sprintf("#  https://docs.fileformat.com/programming/yaml/      #"),
		fmt.Sprintf("#######################################################\n"),
	}, "\n")
}

func (cfg *Config) String() string {
	data, err := os.ReadFile(cfg.location)
	if err != nil {
		return fmt.Sprintf("%v", err.Error())
	}
	return string(data)
}
