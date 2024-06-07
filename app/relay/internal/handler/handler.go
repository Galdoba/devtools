package handler

import (
	"time"

	"github.com/Galdoba/devtools/cronex/job"
)

type handler struct{}

type ExecJob interface {
	TimeValid(time.Time) bool
	Execute() error
}

func Handle(args ...*job.Job) error {
	jobs := make(chan *job.Job)

	go findJobs(jobs, args...)
	for {
		jb, hasMore := <-jobs
		if !hasMore {
			break
		}
		if jb.Halt {
			continue
		}
		if err := jb.Execute(); err != nil {
			return err
		}

	}
	return nil
}

func findJobs(jobs chan *job.Job, inputPool ...*job.Job) {
	now := time.Now()
	for _, jb := range inputPool {
		if jb.TimeValid(now) {
			jobs <- jb
		}
	}
	close(jobs)
}
