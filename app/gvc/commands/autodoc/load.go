package autodoc

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Galdoba/devtools/version"
)

func Load(v *version.Version) (*AutomaticDocumentationFile, error) {
	docPath, err := docPath(v)
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(docPath)
	switch err == nil {
	case false:
		if errors.Is(err, os.ErrNotExist) {
			return new(docPath, v), nil
		}
		return nil, fmt.Errorf("failed to load md file: %v", err)
	case true:
		return load(docPath)

	}
	return nil, fmt.Errorf("not loaded")
}

func load(path string) (*AutomaticDocumentationFile, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v")
	}
	lines := strings.Split(string(bt), "\n")
	amd := &AutomaticDocumentationFile{Path: path, ChangeLog: make(map[versionState][]string)}
	err = amd.parse(lines)
	return amd, err
}

func (amd *AutomaticDocumentationFile) parse(lines []string) error {
	version := versionState{}
	for _, line := range lines {
		switch {
		// case isHeaderLine(line):
		// 	amd.Project = lines[i-1]
		// case isSubheaderLine(line):
		// 	amd.Descr = lines[i-1]
		// case isBoldLine(line):
		// 	amd.OtherHeaderStuff = append(amd.OtherHeaderStuff, line)
		case line == "":
			continue
		case strings.HasPrefix(line, "version "):
			fmt.Println(amd.Path)

			version = parseState(line)
		default:
			amd.ChangeLog[version] = append(amd.ChangeLog[version], line)
		}
	}
	return nil
}

func isHeaderLine(line string) bool {
	for _, s := range strings.Split(line, "") {
		if s != "=" {
			return false
		}
	}
	return true
}

func isSubheaderLine(line string) bool {
	for _, s := range strings.Split(line, "") {
		if s != "-" {
			return false
		}
	}
	return true
}

func isBoldLine(line string) bool {
	if len(line) < 5 {
		return false
	}
	text := strings.Split(line, "")
	for _, i := range []int{0, 1, len(text) - 2, len(text) - 1} {
		if text[i] != "*" {
			return false
		}
	}
	return true
}
