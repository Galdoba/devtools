package cmd

import (
	"fmt"

	"github.com/Galdoba/devtools/app/relay/config"
	"github.com/Galdoba/devtools/cronex/job"
	"github.com/urfave/cli/v2"
)

const ()

func Newjob() *cli.Command {
	cmnd := &cli.Command{
		Name:      "job",
		Aliases:   []string{},
		Usage:     "Create new Job",
		UsageText: "Create new Job in storage directory which will track job fileles and execute them if time is valid",
		//Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
		Action: func(c *cli.Context) error {
			cfg, err := config.Load()
			if err != nil {
				return fmt.Errorf(err.Error())
			}
			handler := c.String("for")
			jobArgs := c.Args().Slice()
			jb := job.Create(handler, jobArgs...)
			jb.Repeatable = c.Bool("repeatable")
			sched := c.String("schedule")
			if sched == "" {
				sched = "* * * * * *"
			}
			jb = jb.SetSchedule(sched)

			return jb.Save(cfg.MessageStorageDirectory())
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "for",
				Usage:    "REQUIRED: utility for the job",
				Required: true,
			},
			&cli.StringFlag{
				Name:        "schedule",
				Usage:       "cron-like time trigger for job",
				DefaultText: "'* * * * *' --> Any time",
				Aliases:     []string{"s"},
			},
			&cli.BoolFlag{
				Name:  "repeatable",
				Usage: "job will be handled every cycle",
			},
		},
	}
	return cmnd
}

/*
relay job -repeatable -schedule "0 * * * *" -for mfline show -s file.mp4
*/
