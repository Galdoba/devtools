package deeptest

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/gconfig"
)

func TestConfig(t *testing.T) {
	cfg, err := gconfig.NewConfig("testProg", gconfig.Default())
	if err != nil {
		fmt.Println("NEW ERR", err.Error())
	}
	//cfg.fillTest()
	//fmt.Println(stdPath("testProg"))
	fmt.Println("NEW", cfg.String())
	cfg.Option_BOOL["main"] = false
	cfg.Option_STR["test3"] = "lskfjaslkgjalkjg"
	if err := cfg.Save(); err != nil {
		fmt.Println("cfg.Save=", err.Error())
	}

	// fmt.Println("SAVE", cfg)
	// cfg2, err2 := gconfig.Load("testProg")

	// //cfg2.SetOptionInt("Test", 2)
	// fmt.Println("LOAD", cfg2, err2)
	// // fmt.Println(cfg2.String("testStr"))
	// // fmt.Println(cfg2.String("Test"))
	// //fmt.Println(cfg2.StringsParameter["testStr"])

	// //log.Fatal("bad config")
	// // //fmt.Println(cfg2.IntFields["Test"] + 3)
	// // cfg2.Save()
	// // fmt.Println(err2)
	// // fmt.Println(cfg2)
	// //fmt.Println(cfg2.IntFields["Test"])

}
