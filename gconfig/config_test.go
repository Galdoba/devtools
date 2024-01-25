package gconfig

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	cfg := newConfig("testProg")
	cfg.fillTest()
	fmt.Println(stdPath("testProg"))
	if err := cfg.Save(); err != nil {
		fmt.Println("cfg.Save=", err.Error())
	}
	cfg2, err2 := Load("testProg")
	cfg2.SetOptionInt("Test", 2)

	fmt.Println(cfg2.IntFields["Test"] + 3)
	cfg2.Save()
	fmt.Println(err2)
	fmt.Println(cfg2)
	fmt.Println(cfg2.IntFields["Test"])
}
