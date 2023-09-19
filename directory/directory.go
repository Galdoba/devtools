package directory

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type content struct {
	dirName string
	name    string
	isDir   bool
}

/*
allFilesPaths := directory.AllFiles(root)
allFilesPaths := directory.List(root, FILES)
allFilesPaths := directory.List(root, FILES_AND_FOLDERS)
allFilesPaths := directory.List(root, FOLDERS)
allFilesPaths := directory.List(root, FOLDERS, CURRENT)
allFilesPaths := directory.List(root, FOLDERS, ALL)
allFilesPaths := directory.List(root, FOLDERS, 1)
allFilesPaths := directory.List(root, directory.ListInstruction{FOLDERS})
allFilesPaths := directory.List(root, instruction.New(FOLDER))


*/

func separator() string {
	return string(filepath.Separator)
}

func Tree(root string) []string {
	list := []string{root}
	for _, leaf := range ListDirs(root) {
		list = append(list, Tree(leaf)...)

	}
	return list
}

func List(dir string) (string, []string, error) {
	path, err := filepath.Abs(dir)
	if err != nil {
		return dir, []string{path, "null", err.Error()}, err
	}
	sep := string(filepath.Separator)
	if !strings.HasSuffix(path, sep) {
		path += sep
	}
	fi, _ := os.ReadDir(dir)
	files := []string{}
	for _, f := range fi {
		files = append(files, f.Name())
	}
	return path, files, nil
}

func ListDirs(root string) []string {
	files := []string{}
	cont, err := contains(root)
	if err != nil {
		return nil
	}
	for _, c := range cont {
		if !c.isDir {
			continue
		}
		newDir := strings.TrimSuffix(c.dirName, separator()) + separator() + c.name + separator()
		files = append(files, newDir)
	}
	return files
}

func ListFiles(root string) []string {
	files := []string{}
	cont, err := contains(root)
	if err != nil {
		return nil
	}
	for _, c := range cont {
		if c.isDir {
			continue
		}
		files = append(files, strings.TrimSuffix(c.dirName, "/")+"/"+c.name)
	}
	return files
}

func ListFilesN(root string, n int) []string {
	files := []string{}
	cont, err := contains(root)
	if err != nil {
		return nil
	}
	for i, c := range cont {
		if i > n {
			break
		}
		if c.isDir {
			continue
		}
		files = append(files, strings.TrimSuffix(c.dirName, "/")+"/"+c.name)
	}
	return files
}

func contains(root string) ([]content, error) {
	root = strings.ReplaceAll(root, "\\", "/")
	var cont []content
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return cont, err
	}
	for _, file := range fileInfo {
		cont = addContent(cont, newContent(root, file.Name(), file.IsDir()))
	}
	return cont, nil
}

func addContent(clist []content, new content) []content {
	clist = append(clist, new)
	return clist
}

func newContent(dirName, name string, isDir bool) content {
	return content{dirName, name, isDir}
}
