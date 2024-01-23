package gconfig

import (
	"encoding/json"
	"fmt"

	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v3"
)

const (
	TYPE_BOOL   = "bool"
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

type config struct {
	Program         string `json "ProgramaaaA"`
	Tags            []string
	loadOrder       []string
	Location        string
	BoolFields      map[string]bool
	IntFields       map[string]int
	FloatFiels      map[string]float64
	StringsFields   map[string]string
	SlBoolFields    map[string][]bool
	SlIntFields     map[string][]int
	SlFloatFiels    map[string][]float64
	SlStringsFields map[string][]string
	keys            map[string]bool

	MapStringsFields map[string]map[string]string
	MapIntFields     map[string]map[string]int
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
	// cfg.mapBoolFields = make(map[string]map[string]bool)
	// cfg.mapIntFields = make(map[string]map[string]int)
	// cfg.mapFloatFiels = make(map[string]map[string]float64)
	cfg.MapStringsFields = make(map[string]map[string]string, 0)
	cfg.MapIntFields = make(map[string]map[string]int, 0)
	return &cfg
}

func (cfg *config) fillTest() {
	cfg.BoolFields[IS_CONCRETE+"bool1"] = true
	cfg.BoolFields[IS_CONCRETE+"bool2"] = true
	cfg.IntFields[IS_CONCRETE+"int1"] = 1
	cfg.IntFields[IS_CONCRETE+"int2"] = 1
	cfg.FloatFiels[IS_CONCRETE+"float1"] = 1.2
	cfg.FloatFiels[IS_CONCRETE+"float2"] = 2.2
	cfg.StringsFields[IS_CONCRETE+"str1_key"] = "str1_val"
	cfg.StringsFields[IS_CONCRETE+"str2_key"] = "str2_val"

	cfg.SlBoolFields[IS_SLICE+"bool1"] = []bool{true, true}
	cfg.SlBoolFields[IS_SLICE+"bool2"] = []bool{true, true}
	cfg.SlIntFields[IS_SLICE+"int1"] = []int{1, 2}
	cfg.SlIntFields[IS_SLICE+"int2"] = []int{3, 4}
	cfg.SlFloatFiels[IS_SLICE+"float1"] = []float64{1.5, 2.5}
	cfg.SlFloatFiels[IS_SLICE+"float2"] = []float64{3.5, 4.5}
	cfg.SlStringsFields[IS_SLICE+"str1_key"] = []string{"str1_val1", "str1_val2"}
	cfg.SlStringsFields[IS_SLICE+"str2_key"] = []string{"str2_val1", "str2_val2"}
	mapFiels1 := make(map[string]string)
	mapFiels1["one"] = "one"
	mapFiels1["two"] = "two"
	mapFiels2 := make(map[string]string)
	mapFiels2["two111"] = "two111"
	mapFiels2["one111"] = "one111"

	mapInt1 := make(map[string]int)
	mapInt1["aaa"] = 1
	mapInt1["bbb"] = 2

	mapInt2 := make(map[string]int)
	mapInt2["ccc"] = 3
	mapInt2["ddd"] = 4

	cfg.MapIntFields["intMap1"] = mapInt1
	cfg.MapIntFields["intMap2"] = mapInt2

	cfg.MapStringsFields["map1"] = mapFiels1
	cfg.MapStringsFields["map2"] = mapFiels2

	bt, err := json.MarshalIndent(cfg, "", "  ")
	fmt.Println(err)
	fmt.Println(string(bt))

	bt2, err2 := toml.Marshal(cfg)
	fmt.Println(err2)
	fmt.Println(string(bt2))

	bt3, err3 := yaml.Marshal(cfg)
	fmt.Println(err3)
	fmt.Println(string(bt3))

}

/*
Path = "c:\\path"
kernels = 5
time_limit = 5.5



*/
