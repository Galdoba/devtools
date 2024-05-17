package main

import (
	"time"

	"github.com/Galdoba/devtools/printer"
	"github.com/Galdoba/devtools/printer/lvl"
)

func main() {
	pm := printer.New().
		WithConsoleColors(true).WithAppName("testAPP").WithFile("tester.log").WithFileLevel(lvl.TRACE).WithConsoleLevel(lvl.INFO)
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		pm.Println(lvl.TRACE, "The answer is 46")
		pm.Println(lvl.INFO, "The answer is 46")

	}

}
