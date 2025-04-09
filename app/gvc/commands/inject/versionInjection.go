package inject

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Galdoba/devtools/version"
)

func Inject(v *version.Version, path string) error {
	lines, err := fileToLines(path)
	if err != nil {
		return err
	}
	vct, err := findVersionControlToken(lines)
	if err != nil {
		return fmt.Errorf("failed to find version control token in %v: %v", path, err)
	}
	newLine, err := inject(v, lines[vct])
	if err != nil {
		return fmt.Errorf("failed to inject version control token in: %v", err)
	}
	lines[vct] = newLine
	if err := updateFile(path, lines); err != nil {
		return fmt.Errorf("failed to update source: %v", err)
	}
	return nil
}

func findVersionControlToken(lines []string) (int, error) {
	candidates := []int{}
	for i, line := range lines {
		lowLine := strings.ToLower(line)
		if strings.Contains(lowLine, " = ") && strings.Contains(lowLine, "//#gvc") {
			candidates = append(candidates, i)
		}
	}
	switch len(candidates) {
	case 0:
		return -1, fmt.Errorf("'#gvc' was not found")
	case 1:
		return candidates[0], nil
	default:
		return -1, fmt.Errorf("multiple version control token's was found")
	}
}

func fileToLines(path string) ([]string, error) {
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}
	lines := strings.Split(string(bt), "\n")
	return lines, nil
}

func inject(v *version.Version, line string) (string, error) {
	re := regexp.MustCompile(`.* = (.*)//#gvc`)
	found := re.FindStringSubmatch(line)
	if len(found) != 2 {
		return "", fmt.Errorf("version control token not found")
	}
	line = strings.ReplaceAll(line, found[1], fmt.Sprintf(`"%v" `, v.String()))
	return line, nil
}

func updateFile(path string, lines []string) error {
	if err := os.Rename(path, path+".tmp"); err != nil {
		return fmt.Errorf("failed to create backup file: %v", err)
	}
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	if err := f.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %v", err)
	}
	text := strings.Join(lines, "\n")
	_, err = f.WriteString(text)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	f.Close()
	os.Remove(path + ".tmp")
	return nil
}
