package autodoc

import (
	"fmt"
	"slices"
)

func (amd *AutomaticDocumentationFile) latest() versionState {
	vs := versionState{}
	keys := []string{}
	for k := range amd.ChangeLog {
		keys = append(keys, k.String())
	}
	switch len(keys) {
	case 0:
	default:
		slices.Sort(keys)
		vs = parseState(keys[len(keys)-1])
	}

	return vs
}

func (amd *AutomaticDocumentationFile) Update(notes ...string) {
	latest := amd.latest()
	for _, note := range notes {
		if note != "" {
			amd.ChangeLog[latest] = append(amd.ChangeLog[latest], fmt.Sprintf("* %v", note))
		}
	}
}

func (amd *AutomaticDocumentationFile) Patch(notes ...string) {
	latest := amd.latest()
	latest.patch++
	for _, note := range notes {
		if note != "" {
			amd.ChangeLog[latest] = append(amd.ChangeLog[latest], fmt.Sprintf("* %v", note))
		}
	}
}

func (amd *AutomaticDocumentationFile) UpdateMinor(notes ...string) {
	latest := amd.latest()
	latest.patch = 0
	latest.minor++
	for _, note := range notes {
		if note != "" {
			amd.ChangeLog[latest] = append(amd.ChangeLog[latest], fmt.Sprintf("* %v", note))
		}
	}
}

func (amd *AutomaticDocumentationFile) UpdateMajor(notes ...string) {
	latest := amd.latest()
	latest.patch = 0
	latest.minor = 0
	latest.major++
	for _, note := range notes {
		if note != "" {
			amd.ChangeLog[latest] = append(amd.ChangeLog[latest], fmt.Sprintf("* %v", note))
		}
	}
}
