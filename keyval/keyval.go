package keyval

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	defaultDir = "/.keyval/"
	key        = 0
	val        = 1
)

func path(name string) string {
	path, _ := os.UserHomeDir()
	path += defaultDir + name
	path = strings.ReplaceAll(path, "\\", "/")
	return path
}

type collection struct {
	path string
	name string
	kval map[string]string
}

func NewCollection(name string) (*collection, error) {
	c := collection{}
	dirname, err := os.UserHomeDir()
	if err != nil {
		return &c, fmt.Errorf("os.UserHomeDir(): %v", err.Error())
	}

	c.path = dirname + defaultDir
	c.path = strings.ReplaceAll(c.path, "\\", "/")
	err = os.MkdirAll(c.path, os.ModePerm)

	if err != nil {
		return &c, fmt.Errorf("os.MkdirAll(%v): %v", c.path, err.Error())
	}
	c.name = name
	pres, err := isPresent(name)
	if pres {
		return nil, fmt.Errorf("can't create: collection is already present")
	}
	if err != nil {
		return nil, fmt.Errorf("can't create: %v", err.Error())
	}

	f, errf := os.Create(c.path + c.name)
	if errf != nil {
		return &c, fmt.Errorf("os.Create(%v): %v", c.path+c.name, errf.Error())
	}
	defer f.Close()
	c.kval = make(map[string]string)
	return &c, nil
}

//separator = :::
func LoadCollection(name string) (*collection, error) {
	pres, err := isPresent(path(name))
	if !pres {
		return nil, fmt.Errorf("can't load collection: collection absent")
	}
	rawData, err := os.ReadFile(path(name))
	if err != nil {
		return nil, fmt.Errorf("can not read: %v", err.Error())
	}
	fmtData := strings.Split(string(rawData[:]), "\n")
	c := collection{}
	c.path = path("")
	c.name = name
	c.kval = make(map[string]string)
	for _, scope := range fmtData {
		data := strings.Split(scope, ":::")
		if len(data) != 2 {
			continue
		}
		c.kval[data[key]] = data[val]
	}
	return &c, nil
}

func SaveCollection(c Kval) error {
	// err := os.Rename(c.Path(), c.Path()+".tmp")

	// bts, err := os.ReadFile(c.Path())
	// if err != nil {
	// 	return fmt.Errorf("original not read: %v", err.Error())
	// }
	// err = os.WriteFile(c.Path()+".tmp", bts, 0600)
	// if err != nil {
	// 	return fmt.Errorf("backup not written: %v", err.Error())
	// }

	f, err := os.OpenFile(c.Path(), os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		return fmt.Errorf("list was not cleared: %v", err.Error())
	}
	defer f.Close()
	f.Truncate(0)
	keys, vals := c.List()
	for i := range keys {
		text := keys[i] + ":::" + vals[i] + "\n"
		if _, err = f.WriteString(text); err != nil {
			return fmt.Errorf("can't write %v | %v", text, err.Error())
		}
	}

	return nil
}

func DeleteCollection(c Kval) error {
	pres, _ := isPresent(c.Path())
	if pres {
		err := os.Remove(c.Path())
		if err != nil {
			return fmt.Errorf("os.Remove(): %v", err.Error())
		}
	}
	c = nil
	return nil
}

type Kval interface {
	Path() string
	List() ([]string, []string)
	Get(string) string
	Set(string, string)
}

func (c *collection) Path() string {
	return c.path + c.name
}

func (c *collection) List() ([]string, []string) {
	ks := []string{}
	for k, _ := range c.kval {
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

///////////////////////////////helpers
func isPresent(name string) (bool, error) {
	if _, err := os.Stat(name); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("file may or may not exist: %v", err.Error())
	}
}
