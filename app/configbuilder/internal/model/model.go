package model

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Galdoba/devtools/csvp"
)

const (
	DataComposition_MAP       = "map"
	DataComposition_SLICE     = "slice"
	DataComposition_PRIMITIVE = "primitive"
	DateType_bool             = "bool"
	DateType_uint8            = "uint8"
	DateType_uint16           = "uint16"
	DateType_uint32           = "uint32"
	DateType_uint64           = "uint64"
	DateType_int8             = "int8"
	DateType_int16            = "int16"
	DateType_int32            = "int32"
	DateType_int64            = "int64"
	DateType_float32          = "float32"
	DateType_float64          = "float64"
	DateType_complex64        = "complex64"
	DateType_complex128       = "complex128"
	DateType_string           = "string"
	DateType_int              = "int"
	DateType_uint             = "uint"
	DateType_uintptr          = "uintptr"
	DateType_byte             = "byte"
	DateType_rune             = "rune"
	EncodingYAML              = "yaml"
	EncodingTOML              = "toml"
	EncodingJSON              = "json"
	EncodingXML               = "xml"
)

type Model struct {
	Fields []*Field
	//rowLen   []int
	language string
}

type Field struct {
	SourceName          string
	DataType            string
	Designation         string
	OmitEmpty           bool
	Comment             string
	DefaulValDictionary map[string]string
	language            string
}

func NewModel(language string) *Model {
	m := Model{}
	m.language = language
	return &m
}

func (m *Model) Delete(i int) error {
	if i < 0 {
		return fmt.Errorf("can't delete: negative index provided")
	}
	if i >= len(m.Fields) {
		return fmt.Errorf("can't delete: no index %v in model fields", i)
	}
	newFields := []*Field{}
	for n, f := range m.Fields {
		if n == i {
			continue
		}
		newFields = append(newFields, f)
	}
	m.Fields = newFields
	return nil
}

func (m *Model) GetFields() []*Field {
	return m.Fields
}

func (m *Model) SwitchFields(a, b int) error {
	if a >= len(m.Fields) || a < 0 {
		return fmt.Errorf("bad index: %v")
	}
	if b >= len(m.Fields) || b < 0 {
		return fmt.Errorf("bad index: %v")
	}
	m.Fields[a], m.Fields[b] = m.Fields[b], m.Fields[a]
	return nil
}

func NewField(language string) *Field {
	f := Field{}
	f.language = language
	return &f
}

func (f *Field) WithSource(source string) *Field {
	f.SourceName = source
	return f
}

func (f *Field) WithDataType(datatype string) *Field {
	f.DataType = datatype
	return f
}

func (f *Field) WithDesignation(designation string) *Field {
	f.Designation = designation
	return f
}
func (f *Field) WithOmitEmpty(oe bool) *Field {
	f.OmitEmpty = oe
	return f
}

func (f *Field) WithComment(comment string) *Field {
	f.Comment = comment
	return f
}

func (f *Field) ToggleOmitempty() *Field {
	f.OmitEmpty = true
	return f
}

func (f *Field) WithValue(key, val string) *Field {
	if f.DefaulValDictionary == nil {
		f.DefaulValDictionary = make(map[string]string)
	}
	if strings.TrimSpace(key) == strings.TrimSpace(val) && strings.TrimSpace(key) == "" {
		return f
	}
	f.DefaulValDictionary[key] = val
	return f
}

func (f *Field) DeleteValue(key string) *Field {
	delete(f.DefaulValDictionary, key)
	return f
}

func (f *Field) Validate() error {
	if len(f.SourceName) == 0 {
		return fmt.Errorf("SourceName must not be blank")
	}
	switch strings.Split(f.SourceName, "")[0] {
	case "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z":
	default:
		return fmt.Errorf("SourceName must start from large latin letter")
	}
	badCharacters := " `!@#$%^&*()-+|'№;%:?	<>" + `"`
	for _, chr := range strings.Split(badCharacters, "") {
		if strings.Contains(f.SourceName, chr) {
			return fmt.Errorf("can't have character '%v' in 'sourcefield'", chr)
		}
	}
	switch f.DataType {
	default:
		if !TypeValid(f.DataType) {
			return fmt.Errorf("unknown primitive '%v' for DataType", f.DataType)
		}
	case "":
		return fmt.Errorf("dataType is not set")
	}
	switch f.Designation {
	case "":
		return fmt.Errorf("designation is not set")
	}
	return nil
}

func (f *Field) String() string {
	oe := ""
	if f.OmitEmpty {
		oe = ",omitempty"
	}
	values := ""
	if f.DefaulValDictionary != nil {
		keys := []string{}
		for k := range f.DefaulValDictionary {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			if k == "" && f.DefaulValDictionary[k] == "" {
				continue
			}
			values += fmt.Sprintf("[%v : %v]", k, f.DefaulValDictionary[k])
		}
	}
	return fmt.Sprintf("%v	%v	`%v:"+`"%v%v"`+"` //%v %v", f.SourceName, f.DataType, f.language, f.Designation, oe, f.Comment, values)
}

func TypeValid(tp string) bool {
	primitives := []string{DateType_bool,
		DateType_uint8, DateType_uint16, DateType_uint32, DateType_uint64,
		DateType_int8, DateType_int16, DateType_int32, DateType_int64,
		DateType_float32, DateType_float64, DateType_complex64, DateType_complex128,
		DateType_string, DateType_int, DateType_uint, DateType_uintptr, DateType_byte, DateType_rune}
	allTypes := append([]string{}, primitives...)
	for _, keys := range primitives {
		allTypes = append(allTypes, "[]"+keys)
		for _, vals := range primitives {
			allTypes = append(allTypes, "map["+keys+"]"+vals)
		}
		for _, vals := range primitives {
			allTypes = append(allTypes, "map["+keys+"][]"+vals)
		}
	}
	for _, t := range allTypes {
		if t == tp {
			return true
		}
	}

	return false
}

// func (m *Model) String() string {
// 	raw := [][]string{}
// 	lenGlob := 0
// 	for _, f := range m.Fields {
// 		rawfld := []string{f.SourceName, f.DataType, f.Designation, fmt.Sprintf("%v", f.OmitEmpty), f.Comment}
// 		for _, k := range keysFrom(f.DefaulValDictionary) {
// 			rawfld = append(rawfld, k)
// 			rawfld = append(rawfld, f.DefaulValDictionary[k])
// 		}
// 		raw = append(raw, rawfld)
// 		if lenGlob < len(rawfld) {
// 			lenGlob = len(rawfld)
// 		}
// 	}
// 	for i, rawOne := range raw {
// 		for len(rawOne) > lenGlob {
// 			rawOne = append(rawOne, " ")
// 		}
// 		for rn, cell := range rawOne {
// 			for len(m.rowLen) <= rn {
// 				m.rowLen = append(m.rowLen, 0)
// 			}
// 			localLettersNum := m.rowLen[rn]
// 			if localLettersNum < len(strings.Split(cell, "")) {
// 				m.rowLen[rn] = len(strings.Split(cell, ""))
// 			}
// 		}
// 		raw[i] = rawOne
// 	}
// 	for i, rawOne := range raw {
// 		for len(rawOne) < lenGlob {
// 			rawOne = append(rawOne, " ")
// 		}
// 		for rn, cell := range rawOne {
// 			for len(strings.Split(cell, "")) < m.rowLen[rn] {
// 				cell += " "
// 			}
// 			rawOne[rn] = cell
// 		}
// 		raw[i] = rawOne
// 	}
// 	lined := []string{}
// 	for _, line := range raw {
// 		lineF := strings.Join(line, `","`)
// 		lineF = `"` + lineF + `"`
// 		lined = append(lined, lineF)

// 	}
// 	lined = append(lined, fmt.Sprintf(`"encoding","%v",""`, m.language))
// 	str := strings.Join(lined, "\n")
// 	return str
// }

func keysFrom(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// func FromString0(s string) (*Model, error) {
// 	lines := strings.Split(s, "\n")
// 	language := strings.Split(lines[len(lines)-1], `","`)[1]
// 	m := NewModel(language)
// 	for i, line := range lines {
// 		if !strings.HasPrefix(line, `"`) || !strings.HasSuffix(line, `"`) {
// 			return nil, fmt.Errorf("line %v does not belong to model", i)
// 		}
// 		line = strings.TrimPrefix(line, `"`)
// 		line = strings.TrimSuffix(line, `"`)
// 		cells := strings.Split(line, `","`)
// 		if i == len(lines)-1 {
// 			m.language = strings.TrimSpace(cells[1])

// 			continue
// 		}

// 		// m.Fields = append(m.Fields, NewField(strings.TrimSpace(cells[0]), strings.TrimSpace(cells[1]), strings.TrimSpace(cells[2])))
// 		m.Fields = append(m.Fields, NewField(language))
// 		// if len(cells) > 5 {
// 		// 	if len(cells)%2 == 0 {
// 		// 		return nil, fmt.Errorf("no dictionary value in line %v", i)
// 		// 	}
// 		// 	m.Fields[i] = m.Fields[i].WithDictionary()
// 		// }
// 		for n, cell := range cells {
// 			cell = strings.TrimSpace(cell)
// 			switch n {
// 			case 0:
// 				m.Fields[i] = m.Fields[i].WithSource(cell)
// 			case 1:
// 				m.Fields[i] = m.Fields[i].WithDataType(cell)
// 			case 2:
// 				m.Fields[i] = m.Fields[i].WithDesignation(cell)
// 			case 3:
// 				om, err := strconv.ParseBool(cell)
// 				if err != nil {
// 					return nil, fmt.Errorf("can't parse cell '%v' (line %v, row %v)", cell, i, n)
// 				}
// 				m.Fields[i] = m.Fields[i].WithOmitEmpty(om)
// 			case 4:
// 				m.Fields[i].Comment = cell
// 			default:
// 				if n%2 == 1 {
// 					continue
// 				}
// 				key := strings.TrimSpace(cells[n-1])
// 				m.Fields[i] = m.Fields[i].WithValue(key, cell)
// 			}
// 		}
// 	}
// 	return m, nil
// }

func FromString(s string) (*Model, error) {
	cntr, err := csvp.FromString(s)
	if err != nil {
		return nil, fmt.Errorf("can't parse string: %v")
	}
	rows := len(cntr.Entries())
	if rows < 2 {
		return nil, fmt.Errorf("need atleast 2 rows to construct a model")
	}
	cols := len(cntr.Entries()[0].Fields())
	if cols < 3 {
		return nil, fmt.Errorf("need atleast 3 columns to construct a model")
	}
	lang := cntr.GetFieldValue(len(cntr.Entries())-1, 1)
	m := NewModel(lang)

	for row, entr := range cntr.Entries() {
		if row == len(cntr.Entries())-1 {
			continue
		}
		modelField := NewField(lang)
		expectValue := false
		for col, entryField := range entr.Fields() {
			switch col {
			case 0:
				modelField.WithSource(entryField)
			case 1:
				modelField.WithDataType(entryField)
			case 2:
				modelField.WithDesignation(entryField)
			case 3:
				oe, err := strconv.ParseBool(entryField)
				if err != nil {
					return nil, fmt.Errorf("can't parse boolean from value (row: %v col: %v val: '%v')", row, col, entryField)
				}
				modelField.WithOmitEmpty(oe)
			case 4:
				modelField.WithComment(entryField)
			default:
				if col%2 == 1 {
					expectValue = true
					continue
				}
				key := strings.TrimSpace(cntr.GetFieldValue(row, col-1))
				modelField.WithValue(key, entryField)
				expectValue = false
			}

		}
		if expectValue {
			return nil, fmt.Errorf("have key column, but no value column")
		}
		m.Fields = append(m.Fields, modelField)
	}
	return m, nil
}

/*
// config struct  
type Config struct {
	AppName           string   `yaml:"App Name                               ,omitempty"`
	Location          string   `yaml:"Location                               ,omitempty"`
	StorageDir        string   `yaml:"Scan Storage Directory                 ,omitempty"`
	WriteLogs         bool     `yaml:"Write Logs                             ,omitempty"`
	LogFile           string   `yaml:"Log File                               ,omitempty"`
	OldScan           float64  `yaml:"Old Scan Age (hours)                   ,omitempty"`
	AutoDeleteOld     bool     `yaml:"Delete Old Scans                       ,omitempty"`
	AutoRenameTracked bool     `yaml:"Change storage data if Tracked renamed ,omitempty"`
	AutoScan          bool     `yaml:"Scan All Files in Tracked Directories  ,omitempty"`
	RescanIfErr       bool     `yaml:"Repeat Scans if Error is met           ,omitempty"`
	TrackDirs         []string `yaml:"Track Directories                      ,omitempty"`
	header            string
}

*/

func FromSource(s string) (*Model, error) {
	validLines := make(map[int]string)
	track := false
	last := 0
	for i, line := range strings.Split(s, "\n") {
		switch line {
		case "type configuration struct {":
			track = true
			continue
		case "}":
			track = false
			continue
		default:
			if !track {
				continue
			}
			validLines[i] = line
			last = i
		}
	}
	m := NewModel("undefined")
	lang := "???"
	for i := 0; i <= last; i++ {
		if _, ok := validLines[i]; !ok {
			continue
		}
		data := strings.Split(validLines[i], "`")
		fld := NewField("undefined")
		fld.DefaulValDictionary = make(map[string]string)
		for j, part := range data {
			switch j {
			case 0:
				structData := strings.Fields(part)
				if len(structData) != 2 {
					return nil, fmt.Errorf("expect 2 struct data fields, have %v %v", len(structData), structData)
				}
				fld.SourceName = strings.TrimSpace(structData[0])
				fld.DataType = strings.TrimSpace(structData[1])
			case 1:
				for _, encodingLang := range []string{EncodingJSON, EncodingTOML, EncodingXML, EncodingYAML} {
					if !strings.Contains(part, encodingLang+":") {
						continue
					}
					fld.language = encodingLang
					lang = encodingLang
					part = strings.TrimPrefix(part, encodingLang+":")
				}
				part = strings.TrimPrefix(part, `"`)
				part = strings.TrimSuffix(part, `"`)
				if strings.HasSuffix(part, ",omitempty") {
					fld.OmitEmpty = true
					part = strings.ReplaceAll(part, ",omitempty", "")
				}
				fld.Designation = strings.TrimSpace(part)
			case 2:
				pairs := strings.Split(part, "[")
				for i, pair := range pairs {
					switch i {
					case 0:
						comments := strings.Split(pair, "//")
						if len(comments) > 0 {
							fld.Comment = comments[1]
						}
					default:
						pair := strings.TrimSuffix(pair, "]")
						kval := strings.Split(pair, " : ")
						if len(kval) != 2 {
							return nil, fmt.Errorf("can't parse key-value pair '%v'", kval)
						}
						fld.DefaulValDictionary[kval[0]] = kval[1]
					}
				}
			}

		}
		if err := fld.Validate(); err != nil {
			fmt.Println(err.Error())
			fmt.Printf("%v : %v", fld.SourceName, err.Error())
			continue
		}
		m.Fields = append(m.Fields, fld)
	}
	m.language = lang
	return m, nil
}

func (m *Model) String() string {
	c := csvp.NewContainer()
	for _, fld := range m.Fields {
		data := []string{}
		data = append(data, fld.SourceName)
		data = append(data, fld.DataType)
		data = append(data, fld.Designation)
		data = append(data, fmt.Sprintf("%v", fld.OmitEmpty))
		data = append(data, fld.Comment)
		keys := keysFrom(fld.DefaulValDictionary)
		for _, k := range keys {
			data = append(data, k)
			data = append(data, fld.DefaulValDictionary[k])
		}
		e := csvp.NewEntry(data...)
		c.AppendEntry(e)

	}
	c.AppendEntry(csvp.NewEntry("encoding", m.language))
	return c.String()
}

func (m *Model) Language() string {
	return m.language
}

func DataTypeSegments(dt string) (string, string, string) {
	composition, keyType, valType := DataComposition_PRIMITIVE, "", ""
	if strings.HasPrefix(dt, "map[") {
		composition = DataComposition_MAP
		dt = strings.TrimPrefix(dt, "map[")
		types := strings.Split(dt, "]")
		keyType = types[0]
		valType = types[1]
		return composition, keyType, valType
	}
	if strings.HasPrefix(dt, "[]") {
		composition = DataComposition_SLICE
		dt = strings.TrimPrefix(dt, "[]")
		valType = dt
		return composition, keyType, valType
	}
	return composition, "", dt
}
