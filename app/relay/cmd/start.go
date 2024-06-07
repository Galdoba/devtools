package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Galdoba/devtools/app/relay/internal/handler"
	"github.com/Galdoba/devtools/cronex/job"
	"github.com/urfave/cli/v2"
)

const ()

func Start() *cli.Command {
	cmnd := &cli.Command{
		Name:      "start",
		Aliases:   []string{},
		Usage:     "Create handler and do Jobs",
		UsageText: "create handler in storage directory which will track job fileles and execute them if time is valid",
		//Description: fmt.Sprintf("build %v file from available model", configbuilder.SOURCE_FILE),
		Action: func(c *cli.Context) error {
			for {
				err := fmt.Errorf("command was not started")
				jobInput := []*job.Job{}
				jobInput, err = collectJobs()
				if err != nil {
					return logger.Errorf("can't collect jobs: %v")
				}
				logger.Println("job input:", len(jobInput))
				jobInput = deleteCompleted(jobInput)
				/*
				   cycle:
				   1 collect jobs:
				   1 delete completed
				   1 filter collected to worker channel
				   1 wait                                  2 do jobs from channel
				*/

				if err := handler.Handle(jobInput...); err != nil {
					return logger.Errorf("%v", err.Error())
				}
				fmt.Println("wait next cycle...")
				time.Sleep(time.Second * time.Duration(cfg.RestartCycle()))
			}
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "fix",
				Usage: "try to fix encountered issues",
			},
		},
	}
	return cmnd
}

func collectJobs() ([]*job.Job, error) {
	jobInput := []*job.Job{}
	fi, err := os.ReadDir(cfg.MessageStorageDirectory())
	if err != nil {
		return nil, logger.Errorf("%v", err.Error())
	}
	for _, f := range fi {
		if f.IsDir() {
			continue
		}
		path := cfg.MessageStorageDirectory() + f.Name()
		jb, err := job.Load(path)
		if err != nil {
			return nil, logger.Errorf("%v", err.Error())
		}
		jobInput = append(jobInput, jb)
	}
	return jobInput, nil
}

func deleteCompleted(jobs []*job.Job) []*job.Job {
	filteredOut := []*job.Job{}
	for _, job := range jobs {
		if job.DeleteImmediatly {
			logger.Println("delete job:", job.ID)
			job.Delete()
			continue
		}
		if !job.Repeatable && job.Done > 0 {
			logger.Println("delete job:", job.ID)
			job.Delete()
			continue
		}
		job.Halt = false
		filteredOut = append(filteredOut, job)
	}
	return filteredOut
}
