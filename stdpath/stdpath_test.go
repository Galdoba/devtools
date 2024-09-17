package stdpath

import (
	"fmt"
	"testing"
)

func TestHome(t *testing.T) {
	home := homeDir()
	fmt.Println(home)
	SetAppName("mgt2")
	fmt.Println(ProgramDir("assets", "characteristic_presets"))
}
