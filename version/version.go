package version

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"

	"github.com/Galdoba/devtools/pathfinder"
)

type Version struct {
	ProjectName  string `json:"name"`
	Description  string `json:"description"`
	MajorVersion int    `json:"major"`
	MinorVersion int    `json:"minor"`
	PatchVersion int    `json:"patch"`
	Build        int    `json:"build"`
	ReleaseDate  string `json:"latest release date,omitempty"` //release = patch/minor/major
	Notes        [][]string
}

func New(options ...VersionOption) *Version {
	v := Version{}
	v.Build = 0
	for _, modifyWithOption := range options {
		modifyWithOption(&v)
	}
	v.Notes = append(v.Notes, []string{})

	return &v
}

type VersionOption func(*Version)

func WithName(name string) VersionOption {
	return func(v *Version) {
		v.ProjectName = name
	}
}

func WithDescription(desc string) VersionOption {
	return func(v *Version) {
		v.Description = desc
	}
}

func (v *Version) Update(notes ...string) {
	v.ReleaseDate = ""
	v.Build++
	v.addNotes(notes...)
}

func (v *Version) Patch(notes ...string) {
	v.ReleaseDate = formatBuildTime()
	v.PatchVersion++
	v.Build++
	v.Notes = append(v.Notes, []string{})
	v.Notes[len(v.Notes)-1] = append(v.Notes[len(v.Notes)-1], fmt.Sprintf("version %v", v.String()))
	v.addNotes(notes...)
}

func (v *Version) UpgradeMinor(notes ...string) {
	v.ReleaseDate = formatBuildTime()
	v.MinorVersion++
	v.PatchVersion = 0
	v.Build++
	v.Notes = append(v.Notes, []string{})
	v.Notes[len(v.Notes)-1] = append(v.Notes[len(v.Notes)-1], fmt.Sprintf("version %v", v.String()))
	v.addNotes(notes...)
}
func (v *Version) UpgradeMajor(notes ...string) {
	v.ReleaseDate = formatBuildTime()
	v.MajorVersion++
	v.MinorVersion = 0
	v.PatchVersion = 0
	v.Build++
	v.Notes = append(v.Notes, []string{})
	v.Notes[len(v.Notes)-1] = append(v.Notes[len(v.Notes)-1], fmt.Sprintf("version %v", v.String()))
	v.addNotes(notes...)
}

func (v *Version) addNotes(notes ...string) {
	for _, note := range notes {
		v.Notes[len(v.Notes)-1] = append(v.Notes[len(v.Notes)-1], note)
	}
}

func (v *Version) String() string {
	s := ""
	s += fmt.Sprintf("%v", v.MajorVersion)
	s += fmt.Sprintf(".%v", v.MinorVersion)
	s += fmt.Sprintf(".%v", v.PatchVersion)
	q := ""

	q = strings.TrimSuffix(q, "-")
	s += q
	if v.ReleaseDate != "" {
		s += ":" + v.ReleaseDate
	}
	s += fmt.Sprintf(" [build %v]", v.Build)
	return s
}

func formatBuildTime() string {
	tm := time.Now()
	mn := int(tm.Month())
	mnth := fmt.Sprintf("%v", mn)
	if len(mnth) < 2 {
		mnth = "0" + mnth
	}
	bt := fmt.Sprintf("%v%v%v", stringed(tm.Year()), stringed(int(tm.Month())), stringed(tm.Day()))
	return bt
}

func stringed(i int) string {
	s := fmt.Sprintf("%v", i)
	for len(s) < 2 {
		s = "0" + s
	}
	return s
}

func Load(name string) (*Version, error) {
	bt, err := os.ReadFile(getpath(name))
	if err != nil {
		return nil, err
	}
	v := Version{}
	if err := json.Unmarshal(bt, &v); err != nil {
		return nil, err
	}
	return &v, nil
}

func getpath(name string) string {
	dir, _ := pathfinder.NewPath(pathfinder.WithProgram("gvc"))
	path := dir + name + ".json"
	return path
}

func (v *Version) Save() error {
	path := getpath(v.ProjectName)
	dir := filepath.Dir(path)
	os.MkdirAll(dir, 0666)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	if err := f.Truncate(0); err != nil {
		return err
	}
	bt, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	_, err = f.Write(bt)
	if err != nil {
		return err
	}
	return f.Close()
}

func (v *Version) Path() string {
	return getpath(v.ProjectName)
}

func (v *Version) Text() string {
	s := ""
	s += fmt.Sprintf("%v\n", v.ProjectName)
	s += fmt.Sprintf("%v\n", v.Description)
	s += fmt.Sprintf(" \n")
	slices.Reverse(v.Notes)
	for _, notes := range v.Notes {
		for _, note := range notes {
			s += fmt.Sprintf("%v\n", note)
		}
		s += fmt.Sprintf(" \n")
	}
	return s
}
