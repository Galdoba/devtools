package autodoc

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/Galdoba/devtools/version"
)

type versionState struct {
	major int
	minor int
	patch int
}

type AutomaticDocumentationFile struct {
	Path             string                    `json:"path"`
	Project          string                    `json:"project"`
	Descr            string                    `json:"description,omitempty"`
	OtherHeaderStuff []string                  `json:"header info"`
	ChangeLog        map[versionState][]string `json:"changelog,omitempty"`
}

func new(path string, v *version.Version) *AutomaticDocumentationFile {
	amd := AutomaticDocumentationFile{}
	amd.Path = path
	amd.Project = v.ProjectName
	amd.Descr = v.Description
	amd.ChangeLog = make(map[versionState][]string)
	amd.OtherHeaderStuff = append(amd.OtherHeaderStuff, fmt.Sprintf("\n**file created at %v**\n", state(v)))
	return &amd
}

/*
gvc
===

description
-----------

**other stuff**

version 1.1.5
* change text 1
* change text 2



*/

func state(v *version.Version) versionState {
	vs := versionState{}
	vs.major = v.MajorVersion
	vs.minor = v.MinorVersion
	vs.patch = vs.patch
	return vs
}

func parseState(s string) versionState {
	fmt.Printf("parse: '%v'\n", s)
	v := versionState{}
	s = strings.TrimPrefix(s, "version")
	mmp := strings.Split(s, ".")
	v.major, _ = strconv.Atoi(mmp[0])
	v.minor, _ = strconv.Atoi(mmp[1])
	v.patch, _ = strconv.Atoi(mmp[2])
	fmt.Printf("got: '%v'\n", v)
	return v
}

func (vs versionState) String() string {
	return fmt.Sprintf("%v.%v.%v", vs.major, vs.minor, vs.patch)
}

func docPath(v *version.Version) (string, error) {
	dir, err := projectRoot(v.Path())
	if err := os.MkdirAll("docs", 0666); err != nil {
		return "", err
	}
	if err != nil {
		return "", err
	}
	return dir + v.ProjectName + ".md", nil
}

func projectRoot(path string) (string, error) {
	dir := filepath.Dir(path)
	projectRootFound := false
	for !projectRootFound {
		fi, err := os.ReadDir(dir)
		if err != nil {
			return "", err
		}
		for _, f := range fi {
			if f.Name() == ".git" && f.IsDir() {
				return fmtDir(dir), nil
			}
		}
		upLayer := filepath.Dir(dir)
		if upLayer == dir {
			break
		}
		dir = upLayer

	}
	return "", fmt.Errorf("no project root was found for %v", path)
}

func fmtDir(dir string) string {
	dir = strings.TrimSuffix(dir, string(filepath.Separator))
	return dir + string(filepath.Separator)
}
