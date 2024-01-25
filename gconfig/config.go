package gconfig

import (
	"fmt"
	"os"
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

type Config interface {
	Path() string
	Data() []byte
}

/*
#######################################################
#  This is auto generated config file.                #
#  Check formatting rules before manual edit.         #
#  https://docs.fileformat.com/programming/yaml/      #
#######################################################

Program: mfline
String Type Parameters:
	default_scan_storage_directory: /home/galdoba/.ffstuff/data/mfline/
	log_file(unimplemented): /home/galdoba/.ffstuff/logs/mfline.log
Float Type Parameters:

*/

// config struct  
type config struct {
	Program          string                        `yaml:"Application Name,omitempty"`
	Tags             []string                      `yaml:"Config tags,omitempty"`
	Location         string                        `yaml:"Config location,omitempty"`
	StringsFields    map[string]string             `yaml:"String Type Parameters,omitempty"`
	IntFields        map[string]int                `yaml:"Integer Type Parameters,omitempty"`
	FloatFiels       map[string]float64            `yaml:"Float Type Parameters,omitempty"`
	BoolFields       map[string]bool               `yaml:"Boolean Type Parameters,omitempty"`
	SlStringsFields  map[string][]string           `yaml:"List of Strings Type Parameters,omitempty"`
	SlIntFields      map[string][]int              `yaml:"List of Integers Type Parameters,omitempty"`
	SlFloatFiels     map[string][]float64          `yaml:"List of Floats Type Parameters,omitempty"`
	SlBoolFields     map[string][]bool             `yaml:"List of Booleans Type Parameters,omitempty"`
	MapStringsFields map[string]map[string]string  `yaml:"Dictionary of Strings Type Parameters,omitempty"`
	MapIntFields     map[string]map[string]int     `yaml:"Dictionary of Integers Type Parameters,omitempty"`
	MapFloatFields   map[string]map[string]float64 `yaml:"Dictionary of Floats Type Parameters,omitempty"`
	MapBoolFields    map[string]map[string]bool    `yaml:"Dictionary of Booleans Type Parameters,omitempty"`
}

func newConfig(program string, tags ...string) *config {
	cfg := config{}
	cfg.Program = program
	for _, t := range tags {
		cfg.Tags = append(cfg.Tags, t)
	}

	cfg.BoolFields = make(map[string]bool)
	cfg.IntFields = make(map[string]int)
	cfg.FloatFiels = make(map[string]float64)
	cfg.StringsFields = make(map[string]string)

	cfg.SlBoolFields = make(map[string][]bool)
	cfg.SlIntFields = make(map[string][]int)
	cfg.SlFloatFiels = make(map[string][]float64)
	cfg.SlStringsFields = make(map[string][]string)

	cfg.MapStringsFields = make(map[string]map[string]string, 0)
	cfg.MapIntFields = make(map[string]map[string]int, 0)
	cfg.MapFloatFields = make(map[string]map[string]float64, 0)
	cfg.MapIntFields = make(map[string]map[string]int, 0)

	return &cfg
}

// fillTest method  
func (cfg *config) fillTest() {
	cfg.StringsFields["log_file(unimplemented)"] = "/home/galdoba/.ffstuff/logs/mfline.log"
	cfg.SetOptionString("scan_storage_directory", "/home/galdoba/.ffstuff/data/mfline/")
	cfg.SetOptionFloat("test", 33.01)
	cfg.SetOptionStringSlice("STR", []string{"aaa", "bbb", "ccc"})
	adMap := make(map[string]string)
	adMap["first"] = "add1"
	adMap["sec"] = "add3"
	adMap["3"] = "add3"
	cfg.SetOptionStringMap("address", adMap)

	bt3, err3 := yaml.Marshal(cfg)
	fmt.Println(err3)
	fmt.Println(string(bt3))

}

func (cfg *config) Save() error {
	bt, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.MkdirAll(strings.TrimSuffix(gpath.StdConfigDir(cfg.Program), "config.yaml"), 0777)
	if err != nil {
		return err
	}
	location := stdPath(cfg.Program)
	f, err := os.OpenFile(location, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	_, err = f.Write(bt)
	return err
}

func Load(program string) (*config, error) {

	bt, err := os.ReadFile(stdPath(program))
	if err != nil {
		return nil, fmt.Errorf("can't load config: %v", err.Error())
	}
	cfg := newConfig(program)
	cfg.SetOptionBool("testing", false)
	err = yaml.Unmarshal(bt, cfg)
	if err != nil {
		return nil, fmt.Errorf("can't load config: %v", err.Error())
	}
	return cfg, nil
}

func (cfg *config) SetOptionString(key, val string) {
	cfg.StringsFields[key] = val
}

func (cfg *config) SetOptionInt(key string, val int) {
	cfg.IntFields[key] = val
}

func (cfg *config) SetOptionFloat(key string, val float64) {
	cfg.FloatFiels[key] = val
}

func (cfg *config) SetOptionBool(key string, val bool) {
	cfg.BoolFields[key] = val
}

func (cfg *config) SetOptionStringSlice(key string, val []string) {
	cfg.SlStringsFields[key] = val
}

func (cfg *config) SetOptionIntSlice(key string, val []int) {
	cfg.SlIntFields[key] = val
}

func (cfg *config) SetOptionFloatSlice(key string, val []float64) {
	cfg.SlFloatFiels[key] = val
}

func (cfg *config) SetOptionBoolSlice(key string, val []bool) {
	cfg.SlBoolFields[key] = val
}

func (cfg *config) SetOptionStringMap(key string, val map[string]string) {
	cfg.MapStringsFields[key] = val
}

func (cfg *config) SetOptionIntMap(key string, val map[string]int) {
	cfg.MapIntFields[key] = val
}

func (cfg *config) SetOptionFloatMap(key string, val map[string]float64) {
	cfg.MapFloatFields[key] = val
}

func (cfg *config) SetOptionBoolMap(key string, val map[string]bool) {
	cfg.MapBoolFields[key] = val
}

func stdPath(program string) string {
	path := gpath.StdConfigDir(program) + "config.yaml"
	return path
}
