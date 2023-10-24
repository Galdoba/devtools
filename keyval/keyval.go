package keyval

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

const (
	key = 0
	val = 1
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

func MakePathJS(path string) string {
	for _, s := range []string{`\`, `/`} {
		pth := strings.Split(path, s)
		if len(pth) > 1 {
			path = strings.ReplaceAll(path, s, sep)
		}
	}
	if strings.HasSuffix(path, ".json") && strings.HasPrefix(path, baseLocation) {
		return path
	}
	dirs := strings.Split(path, sep)
	if dirs[len(dirs)-1] == "" {
		return baseLocation + path
	}
	return baseLocation + path + ".json"
}

type kvalData struct {
	Path   string              `json:"Source"`
	KVpair map[string][]string `json:"Data,omitempty"`
}

func NewKVlist(path string) (*kvalData, error) {
	kv := kvalData{}
	fullpath := MakePathJS(path)
	if !strings.HasSuffix(fullpath, ".json") {
		return nil, fmt.Errorf("can't create: path must not be a directory")
	}

	kv.Path = fullpath
	pres, err := Present(kv.Path)
	if err != nil {
		return nil, fmt.Errorf("can't create: %v", err.Error())
	}
	if pres {
		return nil, fmt.Errorf("can't create: list is already present")
	}

	err = os.MkdirAll(filepath.Dir(kv.Path), os.ModePerm)
	if err != nil {
		return &kv, fmt.Errorf("os.MkdirAll(%v): %v", kv.Path, err.Error())
	}

	f, errf := os.Create(kv.Path)
	if errf != nil {
		return &kv, fmt.Errorf("os.Create(%v): %v", kv.Path, errf.Error())
	}
	defer f.Close()
	kv.KVpair = make(map[string][]string)
	err = kv.Save()
	if err != nil {
		return &kv, err
	}
	return &kv, nil
}

func (kv *kvalData) Save() error {
	//data, err := kv.MarshalJSON()
	fmt.Println(kv.Path)
	data, err := json.MarshalIndent(kv, "", "  ")
	if err != nil {
		return fmt.Errorf("can't save: %v", err.Error())
	}
	f, err := os.OpenFile(kv.Path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("can't save '%v': %v", kv.Path, err.Error())
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("can't save '%v': %v", kv.Path, err.Error())
	}
	if _, err = f.Write(data); err != nil {
		return fmt.Errorf("can't write to '%v': %v", kv.Path, err.Error())
	}
	return nil
}

func Load(name string, flag ...int) (*kvalData, error) {

	fullpath := MakePathJS(name)

	pres, err := Present(fullpath)
	if err != nil {
		return nil, fmt.Errorf("can't confirm collection at '%v'", fullpath)
	}
	if !pres {
		for _, f := range flag {
			if f == os.O_CREATE {
				return NewKVlist(name)
			}
		}
		return nil, fmt.Errorf("no collection at '%v'", fullpath)
	}
	if err := isDirError(fullpath); err != nil {
		return nil, fmt.Errorf("can not read: %v", err.Error())
	}
	rawData, err := os.ReadFile(fullpath)
	if err != nil {
		return nil, fmt.Errorf("can not read: %v", err.Error())
	}
	kv := &kvalData{}
	//	kv.UnmarshalJSON(rawData)
	err = json.Unmarshal(rawData, kv)
	if kv.KVpair == nil {
		kv.KVpair = make(map[string][]string)
	}
	if err != nil {
		return nil, fmt.Errorf("can't Load '%v': %v", name, err.Error())
	}

	return kv, nil
}

func (kv *kvalData) Set(key string, vals ...string) error {
	if len(vals) == 0 {
		return fmt.Errorf("can't set zero values (use Clear() insted)")
	}
	kv.KVpair[key] = vals
	return nil
}

func (kv *kvalData) Add(key string, val string, uniqueOnly bool) error {
	if !uniqueOnly {
		kv.KVpair[key] = append(kv.KVpair[key], val)
		return kv.Save()
	}
	containedVals := kv.KVpair[key]
	for _, check := range containedVals {
		if check == val {
			return nil
		}
	}
	kv.KVpair[key] = append(kv.KVpair[key], val)
	return kv.Save()
}

func (kv *kvalData) UpdateByVal(key string, val string, newVal string) (int, error) {
	values := kv.KVpair[key]
	updated := 0
	for i, present := range values {
		if present == val {
			kv.KVpair[key][i] = newVal
			updated++
		}
	}
	return updated, nil
}

func (kv *kvalData) UpdateByIndex(key string, newVal string, indexes ...int) (int, error) {
	values := kv.KVpair[key]
	updated := 0
	filteredInd := []int{}
	for i, in := range indexes {
		filteredInd = appendUniqueInt(filteredInd, in)
		if i > len(values)-1 {
			return 0, fmt.Errorf("can't update by index: out of bound index provided")
		}
		if in < 0 {
			return 0, fmt.Errorf("can't update by index: negative index provided")
		}
	}
	if len(filteredInd) != len(indexes) {
		return 0, fmt.Errorf("can't update by index: duplicated indexes provided")
	}
	for _, index := range indexes {
		kv.KVpair[key][index] = newVal
		updated++
	}
	return updated, nil
}

func appendUniqueInt(sl []int, i int) []int {
	for _, in := range sl {
		if in == i {
			return sl
		}
	}
	return append(sl, i)

}

func (kv *kvalData) GetSingle(key string) (string, error) {
	vals, ok := kv.KVpair[key]
	switch ok {
	case false:
		return "", fmt.Errorf("no values on key '%s'", key)
	default:
		if len(vals) != 1 {
			return "", fmt.Errorf("muliple values found on key '%s'", key)
		}
		return vals[0], nil
	}
}

func (kv *kvalData) GetAll(key string) ([]string, error) {
	vals, ok := kv.KVpair[key]
	switch ok {
	case false:
		return nil, fmt.Errorf("no values on key '%s'", key)
	default:
		return vals, nil
	}
}

func (kv *kvalData) GetByIndex(key string, indexes ...int) ([]string, error) {
	vals, ok := kv.KVpair[key]
	switch ok {
	case false:
		return nil, fmt.Errorf("no values on key '%s'", key)
	default:
		res := []string{}
		for i, ind := range indexes {
			if ind < 0 {
				return nil, fmt.Errorf("negative index provided (index %d = %d)", i, ind)
			}
			if ind > len(vals)-1 {
				return nil, fmt.Errorf("out of bound index provided (index %d = %d)", i, ind)
			}
			res = append(res, vals[ind])
		}
		return res, nil
	}
}

func (kv *kvalData) RemoveByVal(key string, vals ...string) error {
	keep := []string{}
	have := kv.KVpair[key]
	if len(vals) == 0 {
		return fmt.Errorf("can't remove: input have zero values")
	}
	if len(have) == 0 {
		return fmt.Errorf("can't remove: nothing to remove from")
	}
mLoop:
	for _, present := range have {
		for _, check := range vals {
			if present == check {
				continue mLoop
			}
		}
		keep = append(keep, present)
	}
	kv.KVpair[key] = keep
	return nil
}

func (kv *kvalData) RemoveByKey(key string) error {
	_, err := kv.GetAll(key)
	if err != nil {
		return err
	}
	delete(kv.KVpair, key)
	return nil
}

func Delete(kv *kvalData) error {
	path := kv.Path
	js, err := os.Stat(path)
	if js.IsDir() {
		return fmt.Errorf("can't delete: %v", err)
	}
	if err != nil {
		return fmt.Errorf("can't delete: %v", err)
	}
	kv = nil
	return os.Remove(path)
}

func appendUnique(sl []string, s string) []string {
	for _, str := range sl {
		if str == s {
			return sl
		}
	}
	return append(sl, s)
}

func Present(name string) (bool, error) {
	path := MakePathJS(name)
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("file may or may not exist: %v", err.Error())
	}
}

func KVlistPresent(path string) bool {
	switch strings.HasSuffix(path, ".json") {
	case true:
		f, err := os.Stat(path)
		if err != nil {
			return false
		}
		if f.IsDir() {
			return false
		}
		return true
	default:
		path = MakePathJS(path)
		_, err := os.Stat(path)
		if err != nil {
			return false
		}
		return true
	}
}

func (kv *kvalData) Data() map[string][]string {
	return kv.KVpair
}

func (kv *kvalData) Keys() []string {
	keys := []string{}
	for k := range kv.KVpair {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
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
