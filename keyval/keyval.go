package keyval

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"

	"github.com/Galdoba/devtools/directory"
)

const (
	KVRecordSep = "<=}\n"
	KVUnitSep   = "{=>"
	key         = 0
	val         = 1
)

var baseLocation string
var sep string

func init() {
	osname := runtime.GOOS
	switch osname {
	default:
		panic(fmt.Sprintf("'%v' operating system is not supported", osname))
	case "windows", "linux":
		sep = string(filepath.Separator)
	}
	userDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("can't get userDir: %v", err.Error()))
	}
	baseLocation = userDir + sep + ".keyval" + sep
}

// MakePath - assemble path for collection with input data
func MakePath(path string) string {
	if strings.HasSuffix(path, ".kv") && strings.HasPrefix(path, baseLocation) {
		return path
	}
	dirs := strings.Split(path, sep)
	if dirs[len(dirs)-1] == "" {
		return baseLocation + path
	}
	return baseLocation + path + ".kv"
}

func MapCollection(path string) map[string]string {
	path = MakePath(path)
	kvmap := make(map[string]string)
	files := directory.Tree(path)
	for _, fl := range files {
		if !strings.HasSuffix(fl, ".kv") {
			continue
		}
		rawData, _ := os.ReadFile(fl)
		fmtData := strings.Split(string(rawData[:]), "\n")
		for _, scope := range fmtData {
			data := strings.Split(scope, KVUnitSep)
			if data[key] == "" {
				continue
			}
			value := strings.Join(data[val:], KVUnitSep)
			kvmap[data[key]] = cleanCrLf(value)
		}
	}

	return kvmap
}

type collection struct {
	path string
	kval map[string]string
}

func NewCollection(path string) (*collection, error) {
	c := collection{}
	fullpath := MakePath(path)
	if !strings.HasSuffix(fullpath, ".kv") {
		return nil, fmt.Errorf("can't create: path must not be a directory")
	}

	c.path = fullpath

	pres, err := Present(c.path)
	if err != nil {
		return nil, fmt.Errorf("can't create: %v", err.Error())
	}
	if pres {
		return nil, fmt.Errorf("can't create: collection is already present")
	}

	err = os.MkdirAll(filepath.Dir(c.path), os.ModePerm)
	if err != nil {
		return &c, fmt.Errorf("os.MkdirAll(%v): %v", c.path, err.Error())
	}

	f, errf := os.Create(c.path)
	if errf != nil {
		return &c, fmt.Errorf("os.Create(%v): %v", c.path, errf.Error())
	}
	defer f.Close()
	c.kval = make(map[string]string)
	return &c, nil
}

func SliceValues(str string) []string {
	return strings.Split(str, KVUnitSep)
}

// separator = :::    key1{=>val1{=>val2<=}\n
//
//	|=>
//
// <===>
// <=|/n
func LoadCollection(path string) (*collection, error) {
	fullpath := MakePath(path)

	pres, err := Present(fullpath)
	if err != nil {
		return nil, fmt.Errorf("can't confirm collection at '%v'", path)
	}
	if !pres {
		return nil, fmt.Errorf("no collection at '%v'", path)
	}
	if err := isDirError(fullpath); err != nil {
		return nil, fmt.Errorf("can not read: %v", err.Error())
	}
	rawData, err := os.ReadFile(fullpath)
	if err != nil {
		return nil, fmt.Errorf("can not read: %v", err.Error())
	}
	fmtData := strings.Split(string(rawData[:]), KVRecordSep)
	c := collection{}
	c.path = fullpath
	c.kval = make(map[string]string)
	for _, scope := range fmtData {
		data := strings.Split(scope, KVUnitSep)
		value := strings.Join(data[val:], KVUnitSep)
		c.kval[data[key]] = cleanCrLf(value)
	}
	return &c, nil
}

func cleanCrLf(s string) string {
	s = strings.TrimSuffix(s, "\n")
	s = strings.TrimSuffix(s, "\r")
	return s
}

func SaveCollection(c Kval) error {
	f, err := os.OpenFile(c.Path(), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("can't save '%v': %v", c.Path(), err.Error())
	}
	defer f.Close()
	keys, vals := c.List()
	text := ""
	for i, k := range keys {
		if k == "" {
			continue
		}
		text += keys[i] + KVUnitSep + vals[i] + KVRecordSep
	}
	if text == "" {
		return nil
	}
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("can't save '%v': %v", c.Path(), err.Error())
	}
	if _, err = f.WriteString(text); err != nil {
		return fmt.Errorf("can't write to '%v': %v", c.Path(), err.Error())
	}
	return nil
}

func DeleteCollection(name string) error {
	path := MakePath(name)
	if err := isDirError(path); err != nil {
		return fmt.Errorf("can't delete: %v", err.Error())
	}
	if err := os.Remove(path); err != nil {
		return fmt.Errorf("can't delete: %v", err.Error())
	}
	return nil
}

type Kval interface {
	Path() string
	List() ([]string, []string)
	Get(string) string
	Set(string, string)
	Destroy() error
}

func (c *collection) Path() string {
	return c.path
}

func (c *collection) List() ([]string, []string) {
	ks := []string{}
	for k := range c.kval {
		if k == "" {
			continue
		}
		ks = append(ks, k)
	}
	sort.Strings(ks)
	vls := []string{}
	for _, k := range ks {
		vls = append(vls, c.kval[k])
	}
	return ks, vls
}

func (c *collection) Get(key string) string {
	if v, ok := c.kval[key]; ok {
		return v
	}
	return "[NULL]"
}

func (c *collection) Set(key, val string) {
	c.kval[key] = val
}

func (c *collection) Clear(key string) {
	delete(c.kval, key)
}

func (c *collection) Destroy() error {
	return os.Remove(c.Path())
}

// /////////////////////////////helpers
func Present(name string) (bool, error) {
	path := MakePath(name)
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("file may or may not exist: %v", err.Error())
	}
}

func isDirError(path string) error {
	fi, err := os.Stat(path)
	if err != nil {
		return err
	}
	if fi.IsDir() {
		return fmt.Errorf("%v is directory", path)
	}
	return nil
}
