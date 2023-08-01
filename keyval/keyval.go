package keyval

import (
	"fmt"
	"os"
)

type collection struct {
	path string
	kval map[string]string
	sep  string
}

type Kval interface {
	Path() string
	Keys(string) []string
	Call(string) string
}

func New(name string) (*collection, error) {
	c := collection{}
	dirname, err := os.UserHomeDir()
	if err != nil {
		return &c, fmt.Errorf("os.UserHomeDir(): %v", err.Error())
	}
	panic(dirname)
	c.path = dirname
	return &c, nil
}

/*
kval.Call("base_name_s01")
зуьфдентщм
*/
