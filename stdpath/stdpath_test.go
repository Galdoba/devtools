package stdpath

import (
	"fmt"
	"testing"
)

func TestHome(t *testing.T) {

	SetAppName("my_app")
	fmt.Println("config dir              :", ConfigDir())
	fmt.Println("config file             :", ConfigFile())
	fmt.Println("config file with key    :", ConfigFile("custum_1"))
	fmt.Println("log dir                 :", LogDir())
	fmt.Println("log file                :", LogFile())
	fmt.Println("log file with key       :", LogFile("custom_2"))
	fmt.Println("program dir             :", ProgramDir())
	fmt.Println("program dir with layers :", ProgramDir("layer1", "layer2"))
}
