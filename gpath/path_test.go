package gpath

import "testing"

func pathInput() []string {
	return []string{
		"input",
		`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\gpath\path.go`,
		`d:\MUX\MUX_2022-03-21\`,
		"d:\\MUX\\MUX_2022-03-25\\",
	}
}

func TestPath(t *testing.T) {
	for num, input := range pathInput() {
		n := num + 1
		p := newPath(input)
		if p == nil {
			t.Errorf("path is <nil>")
		}
		if p.input == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.input is not set", n, input)
		}
		if p.server == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.server is not set", n, input)
		}
		if p.tome == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.tome is not set", n, input)
		}
		if p.dir == nil {
			t.Errorf("Test %v:\ninput: '%v'\n	p.dir is not set", n, input)
		}
		if p.fileName == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.fileName is not set", n, input)
		}
		if p.extention == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.extention is not set", n, input)
		}
		if p.osyst == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.osyst is not set", n, input)
		}
		if p.pathType == UNSPECIFIED {
			t.Errorf("Test %v:\ninput: '%v'\n	p.pathType is not set", n, input)
		}
		if p.pathType != SourcePath && p.pathType != DestinationPath && p.pathType != DataPath && p.pathType != UNSPECIFIED {
			t.Errorf("Test %v:\ninput: '%v'\n	p.pathType is unknown", n, input)
		}
	}
}
