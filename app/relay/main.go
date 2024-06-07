package main

import (
	"fmt"
	"os"

	"github.com/Galdoba/devtools/app/relay/cmd"
	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/configmanager"
	"github.com/Galdoba/devtools/gpath"
	"github.com/urfave/cli/v2"
)

const (
	programName = "relay"
)

func setupConfig() error {
	cfg := config.New()
	if err := cfg.SetDefault(); err != nil {
		fmt.Printf("config setup failed: %v\n", err.Error())
		os.Exit(1)
	}
	cfg.SetMessageStorageDirectory(gpath.AppStorageDir(programName))
	if err := cfg.Save(); err != nil {
		fmt.Printf("config setup failed: %v\n", err.Error())
		os.Exit(1)
	}
	return nil
}

func main() {
	fmt.Println("start main")
	app := cli.NewApp()
	app.Version = "v 0.2.0"
	app.Name = programName
	app.Usage = "Track command queue and relay operations to other apps"
	app.Description = "as a daemon track particular dir and run commands if all is good"
	app.Flags = []cli.Flag{}
	// An action to execute before any subcommands are run, but after the context is ready
	// If a non-nil error is returned, no subcommands are run
	cfgPath, err := configmanager.DefaultConfigPath(programName)
	if cfgPath == "" {
		err = setupConfig()
	}
	fmt.Println(cfgPath, "---")
	// _, err = config.Load()
	// if err != nil {
	// 	fmt.Println("no config")
	// 	cfg := config.New()
	// 	if err := cfg.SetDefault(); err != nil {
	// 		fmt.Println("default config setup failed")
	// 		os.Exit(1)
	// 	}
	// 	cfg.SetMessageStorageDirectory(gpath.AppStorageDir(app.Name))
	// 	if err := cfg.Save(); err != nil {
	// 		fmt.Println("default config write failed")
	// 		os.Exit(1)
	// 	}
	// 	fmt.Println("default config setup complete")
	// 	fmt.Printf("check %v\n", gpath.StdConfigDir(app.Name))

	// }
	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []*cli.Command{
		cmd.Health(),
		cmd.Newjob(),
		cmd.Start(),
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err = app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}
