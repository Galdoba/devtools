package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/cls"
)

var cfg config.Config
var logger = cls.New()

func init() {
	err := fmt.Errorf("no init made")
	cfg, err = config.Load()
	if err != nil {
		fmt.Printf("can't init config: %v", err.Error())
		os.Exit(1)
	}
	if err = logger.AddFile(cfg.LogLocation(), cfg.LoggerLevel(), "", log.Ldate|log.Ltime); err != nil {
		fmt.Printf("can't setup logger: \n%v", err.Error())
		os.Exit(1)
	}
	_, err = os.ReadDir(cfg.MessageStorageDirectory())
	if err != nil {
		fmt.Printf("can't read Storage Directory (%v): \n%v", cfg.MessageStorageDirectory(), err.Error())
		os.Exit(1)
	}
}
