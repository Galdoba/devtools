package autodoc

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func (amd *AutomaticDocumentationFile) Save() error {
	if amd.Path == "" {
		return fmt.Errorf("no path to save amd")
	}
	f, err := os.Create(amd.Path)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	text := amd.ToText()
	f.Truncate(0)
	_, err = f.WriteString(text)
	return f.Close()
}

func (amd *AutomaticDocumentationFile) ToText() string {
	lines := []string{}

	keyMap := make(map[string]versionState)
	for k := range amd.ChangeLog {
		keyMap[k.String()] = k
	}
	keys := []string{}
	for k := range keyMap {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	for _, kStr := range keys {
		version := parseState(kStr)
		lines = append(lines, fmt.Sprintf("version %v", version.String()))
		for _, note := range amd.ChangeLog[version] {
			lines = append(lines, note)
		}
		lines = append(lines, "\n")
	}
	return strings.Join(lines, "\n")
}
