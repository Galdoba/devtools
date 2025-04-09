package autodoc

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
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
	path      string
	Project   string `json:"project"`
	Descr     string `json:"description,omitempty"`
	version   *version.Version
	ChangeLog map[versionState][]string `json:"changelog,omitempty"`
}

func new(path string, v *version.Version) *AutomaticDocumentationFile {
	amd := AutomaticDocumentationFile{}
	amd.path = path
	amd.version = v
	amd.Project = v.ProjectName
	amd.Descr = v.Description
	amd.ChangeLog = make(map[versionState][]string)
	return &amd
}

func stateIsLater(main, compareWith versionState) bool {
	if main.major > compareWith.major {
		return true
	}
	if main.minor > compareWith.minor {
		return true
	}
	if main.patch > compareWith.patch {
		return true
	}
	return false
}

func (amd *AutomaticDocumentationFile) statesbackward() []versionState {
	stateMap := make(map[int]versionState)
	keys := []int{}
	for state := range amd.ChangeLog {
		stateMap[state.val()] = state
		keys = append(keys, state.val())
	}
	vals := []versionState{}
	slices.Sort(keys)
	slices.Reverse(keys)
	for _, k := range keys {
		vals = append(vals, stateMap[k])
	}
	return vals
}

func (vs versionState) val() int {
	return vs.patch + (vs.minor * 100) + (vs.major * 10000)
}

func (amd *AutomaticDocumentationFile) Text() string {
	s := ""
	s += fmt.Sprintf("%v\n", amd.Project)
	s += fmt.Sprintf("%v\n", amd.Descr)
	s += fmt.Sprintf(" \n")
	s += fmt.Sprintf("Change Log:\n")
	for _, state := range amd.statesbackward() {
		s += fmt.Sprintf(" version %v\n", state.String())
		for _, change := range amd.ChangeLog[state] {
			s += fmt.Sprintf("  -%v\n", change)
		}
		s += fmt.Sprintf(" \n")
	}
	return s
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

// func parseState(s string) versionState {
// 	fmt.Printf("parse: '%v'\n", s)
// 	v := versionState{}
// 	s = strings.TrimPrefix(s, "version")
// 	mmp := strings.Split(s, ".")
// 	v.major, _ = strconv.Atoi(mmp[0])
// 	v.minor, _ = strconv.Atoi(mmp[1])
// 	v.patch, _ = strconv.Atoi(mmp[2])
// 	fmt.Printf("got: '%v'\n", v)
// 	return v
// }

func stateFromString(s string) (versionState, error) {
	re := regexp.MustCompile(`(\d+)\.(\d+)\.(\d+)`)
	subs := re.FindStringSubmatch(s)
	switch len(subs) {
	case 4:
		vs := versionState{}
		vs.major, _ = strconv.Atoi(subs[1])
		vs.minor, _ = strconv.Atoi(subs[2])
		vs.patch, _ = strconv.Atoi(subs[3])
		return vs, nil
	}
	return versionState{}, fmt.Errorf("failed to parse '%v' to version state", s)
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
