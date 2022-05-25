package directory

import (
	"io/ioutil"
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

func Tree(root string) []string {
	list := []string{root}
	for _, leaf := range ListDirs(root) {
		list = append(list, Tree(leaf)...)
	}
	return list
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
		newDir := strings.TrimSuffix(c.dirName, "/") + "/" + c.name
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
