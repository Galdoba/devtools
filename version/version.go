package version

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Version struct {
	path         string
	MajorVersion int      `json:"major"`
	MinorVersion int      `json:"minor"`
	PatchVersion int      `json:"patch"`
	Qualifiers   []string `json:"qualifiers,omitempty"`
	Build        int      `json:"build"`
	BuildDate    string   `json:"latest build date,omitempty"`
}

func New(path string) *Version {
	v := Version{}
	v.path = path
	v.Build = 1
	return &v
}

func (v *Version) Update(qual ...string) {
	v.BuildDate = ""
	v.Qualifiers = qual
	v.Build++
}

func (v *Version) Patch(qual ...string) {
	v.BuildDate = formatBuildTime()
	v.Qualifiers = qual
	v.PatchVersion++
	v.Build++
}

func (v *Version) UpgradeMinor(qual ...string) {
	v.BuildDate = formatBuildTime()
	v.Qualifiers = qual
	v.MinorVersion++
	v.PatchVersion = 0
	v.Build++
}
func (v *Version) UpgradeMajor(qual ...string) {
	v.BuildDate = formatBuildTime()
	v.Qualifiers = qual
	v.MajorVersion++
	v.MinorVersion = 0
	v.PatchVersion = 0
	v.Build++
}

func (v *Version) String() string {
	s := ""
	s += fmt.Sprintf("%v", v.MajorVersion)
	s += fmt.Sprintf(".%v", v.MinorVersion)
	s += fmt.Sprintf(".%v", v.PatchVersion)
	q := ""
	if len(v.Qualifiers) > 0 {

		for _, qual := range v.Qualifiers {
			q += "-" + qual
		}

	}
	q = strings.TrimSuffix(q, "-")
	s += q
	if v.BuildDate != "" {
		s += ":" + v.BuildDate
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

func Load(path string) (*Version, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	v := Version{}
	if err := json.Unmarshal(bt, &v); err != nil {
		return nil, err
	}
	v.path = path
	return &v, nil
}

func (v *Version) Save() error {
	f, err := os.OpenFile(v.path, os.O_CREATE|os.O_WRONLY, 0666)
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
