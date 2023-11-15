package gpath

import (
	"fmt"
	"testing"
)

func pathInput() []string {
	return []string{
		"input",
		`c:\Users\pemaltynov\go\src\github.com\Galdoba\devtools\gpath\path.go`,
		`d:\MUX\MUX_2022-03-21\`,
		`.\MUX\15-15`,
		"d:\\MUX\\MUX_2022-03-25\\",
		"\\\\192.168.31.4\\root\\EDIT\\@trailers_temp\\Last_Survivors_TRL_AUDIORUS20.m4a",
		``,
		"",
		".",
		"f.a",
		"d:\\IN\\IN_2022-03-28\\LINK_to_File.lnk",
		"d:\\IN\\IN_2022-03-28\\LINK_to_Dir.lnk",
	}
}

func TestPath(t *testing.T) {
	// if err := Touch(`c:\Users\pemaltynov\tabviewer\DataFile.csv`); err != nil {
	// 	t.Errorf(err.Error())
	// }
	// fmt.Println(StdPath("data.txt", []string{"category", ".hidden", ".config", "program"}...))

	return
	for num, input := range pathInput() {
		n := num + 1
		fmt.Println("Start test: '" + input + "'")
		p, err := newPath(input)
		if err != nil {
			t.Errorf("internal error: %v", err.Error())
			fmt.Println(" ")
			continue
		}
		if p == nil {
			t.Errorf("path is <nil>")
			continue
		}
		if p.input == "" {
			t.Errorf("Test %v:\ninput: '%v'\n	p.input is not set", n, input)
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
		fmt.Println("SUCCESS!!")
	}
}
