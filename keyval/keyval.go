package keyval

import (
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

const (
	defaultDir = "/.keyval/"
	key        = 0
	val        = 1
)

func path(name string) string {
	path, _ := os.UserHomeDir()
	path += defaultDir + name
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

	_, err = os.Create(c.path + c.name)
	if err != nil {
		return &c, fmt.Errorf("os.Create(%v): %v", c.path+c.name, err.Error())
	}
	c.kval = make(map[string]string)
	return &c, nil
}

//separator = :::
func LoadCollection(name string) (*collection, error) {
	pres, err := isPresent(name)
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
	err := os.Rename(c.Path(), c.Path()+".tmp")
	i := 0
	for err != nil && i < 100 {
		fmt.Printf("Try save %v/100\r", i)
		err = os.Rename(c.Path(), c.Path()+".tmp")
		time.Sleep(time.Millisecond * 10)
		i++
	}
	if err != nil {
		return fmt.Errorf("backup not created: %v", err.Error())
	}
	f, err := os.OpenFile(c.Path(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("list was not cleared: %v", err.Error())
	}
	defer f.Close()

	keys, vals := c.List()
	for i := range keys {
		text := keys[i] + ":::" + vals[i] + "\n"
		if _, err = f.WriteString(text); err != nil {
			return fmt.Errorf("can't write %v | %v", text, err.Error())
		}
	}
	err = os.Remove(c.Path() + ".tmp")
	if err != nil {
		return fmt.Errorf("os.Remove() backup: %v", err.Error())
	}
	return nil
}

func DeleteCollection(c Kval) error {

	err := os.Remove(c.Path())
	if err != nil {
		return fmt.Errorf("os.Remove(): %v", err.Error())
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
	dirname, _ := os.UserHomeDir()
	path := dirname + defaultDir

	if _, err := os.Stat(path + name); err == nil {
		return true, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else {
		return false, fmt.Errorf("file may or may not exist: %v", err.Error())
	}
}
