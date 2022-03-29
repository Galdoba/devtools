package gpath

import (
	"path/filepath"
	"strings"
)

/*
TODO:
создать инструментарий позволяющий манипулировать файлами и ориентацией в пространстве для сторонних утилит

*/

const (
	UNSPECIFIED = iota
	SourcePath
	DestinationPath
	DataPath
)

type path struct {
	input     string
	server    string
	tome      string
	dir       []string
	fileName  string
	extention string
	osyst     string
	pathType  int
}

func newPath(input string, pathType ...int) *path {
	p := path{}
	p.input = input
	if !strings.Contains(input, `\\`) {
		filepath.Abs()
	}
	return &p
}
