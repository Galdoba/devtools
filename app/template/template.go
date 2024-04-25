package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

/*
run


*/

var configPath string

const (
	programName = "app_name"
)

func main() {

	app := cli.NewApp()
	app.Version = "v 0.0.1"
	app.Name = programName
	app.Usage = "what do i do?"
	app.Flags = []cli.Flag{}
	//�� ������ ��������
	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.Commands = []*cli.Command{

		// cmd.NewCommand(),
	}

	//�� ��������� ��������
	app.After = func(c *cli.Context) error {
		return nil
	}
	args := os.Args
	if err := app.Run(args); err != nil {
		errOut := fmt.Sprintf("%v error: %v", programName, err.Error())
		println(errOut)
		os.Exit(1)
	}

}
