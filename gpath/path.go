package gpath

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

type pathStr struct {
	input      string
	executable string
	computer   string
	tome       string
	dir        []string
	fileName   string
	extention  string
	osyst      string
	pathType   int
	isDir      bool
	isFile     bool
	isAbs      bool
}

func newPath(input string) (*pathStr, error) {
	if strings.TrimSpace(input) == "" {
		return nil, fmt.Errorf("path was not specified")
	}
	abs, err := filepath.Abs(input)
	if err != nil {
		return nil, err
	}
	pathStats, errStat := os.Stat(input)
	if errStat != nil {
		return nil, errStat
	}
	p := pathStr{}
	p.input = input
	if abs == input {
		p.isAbs = true
	}
	p.isDir = pathStats.Mode().IsDir()
	p.isFile = pathStats.Mode().IsRegular()
	p.executable, _ = os.Executable()
	p.osyst = runtime.GOOS
	slashedInput := filepath.ToSlash(input)
	data := strings.Split(slashedInput, "/")
	p.tome = filepath.VolumeName(slashedInput)
	if p.isDir {
		p.extention = ""
	}
	if strings.Contains(slashedInput, "/") {
		comp := strings.Split(slashedInput, "/")
		compName := strings.Split(comp[1], "/")
		p.computer = compName[0]
	}
	p.fileName = filepath.Base(slashedInput)

	fn := strings.Split(p.fileName, ".")
	if len(fn) > 1 {
		p.extention = fn[len(fn)-1]
	}
	for _, dr := range data {
		if dr == "" || dr == p.computer || dr == p.fileName || dr == p.tome {
			continue
		}
		p.dir = append(p.dir, dr)
	}
	if p.tome == "" {
		return &p, fmt.Errorf("tome is not set [%v]", input)
	}
	if p.dir == nil {
		return &p, fmt.Errorf("dir is not set [%v]", input)
	}
	if p.fileName == "" {
		return &p, fmt.Errorf("fileName is not set [%v]", input)
	}
	if p.extention == "" {
		return &p, fmt.Errorf("extention is not set [%v]", input)
	}
	if p.osyst == "" {
		return &p, fmt.Errorf("osyst is not set [%v]", input)
	}
	return &p, nil
}

func Touch(path string) error {
	if _, err := os.Stat(path); err == nil {
		return nil
	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		dir := filepath.Dir(path)
		if err := os.MkdirAll(dir, 0777); err != nil {
			return err
		}
		f, err := os.OpenFile(path, os.O_CREATE, 0777)
		if err != nil {
			return err
		}
		return f.Close()
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		return err
	}
}

func StdPath(name string, dirs ...string) string {
	path, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	sep := string(filepath.Separator)
	path += sep
	for _, dir := range dirs {
		path += dir + sep
	}
	return path + name

}
