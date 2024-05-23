package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/Galdoba/devtools/app/tabledit/config"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/urfave/cli/v2"
)

var cfg config.Config

func Open() *cli.Command {
	cmnd := &cli.Command{
		Name:      "open",
		Aliases:   []string{},
		Usage:     "Open Table",
		UsageText: "tabledit open [command options]",
		Description: strings.Join([]string{
			"Interactive loop creates table model and allows to edit it",
		}, "\n"),
		Action: func(c *cli.Context) error {
			cnfg, err := config.Load()
			if err != nil {
				return fmt.Errorf("can't start command: %v", err.Error())
			}
			cfg = cnfg
			//Setup Logger
			f, err := tea.LogToFile(cfg.LogPath(), "debug")
			if err != nil {
				log.Fatalf("LogToFile: %v", err.Error())
			}
			defer f.Close()
			//Create table model
			m, err := NewModel()
			if err != nil {
				return fmt.Errorf("can't create model: %v", err.Error())
			}
			//Initiate Buble-Tea
			p := tea.NewProgram(m, tea.WithAltScreen())
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}

			//Edit Loop

			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "overwrite",
				Usage:   "allow overwrite of existing model's file",
				Aliases: []string{"o"},
			},
		},
	}

	return cmnd

}
