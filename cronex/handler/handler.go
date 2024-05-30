package handler

import (
	"fmt"
	"os"
	"time"

	"github.com/Galdoba/devtools/cronex/job"
	"github.com/Galdoba/devtools/cronex/job/schedule"
)

type Option struct {
	cycle int
}

func Cycle(i int) Option {
	return Option{
		cycle: i,
	}
}

type Handler interface {
	Asses(*job.Job) string
	Handle(string, *job.Job) error
}

type handler struct {
	storage string
	cycle   int
}

func New(path string) *handler {
	h := handler{}
	h.storage = path
	return &h
}

func (h *handler) With(options ...Option) *handler {
	for _, o := range options {
		if o.cycle > 0 {
			h.cycle = o.cycle
		}
	}
	return h
}

func (h *handler) Asses(j *job.Job) string {
	next := "skip"
	if j.DeleteImmediatly {
		return "delete"
	}
	if !j.Repeatable && j.Done > 0 {
		return "skip"
	}
	if schedule.TimeValid(j.Schedule) {
		next = "act"
		if j.Skip {
			next = "skip"
		}
	}
	return next
}

func (h *handler) Handle(j *job.Job) error {
	action := h.Asses(j)
	for {
		switch action {
		case "wait":
			panic("----")
		case "act":
			action = "done"
			out := ""
			err := fmt.Errorf("no exec")
			if out, err = j.Execute(); err != nil {
				if j.DeleteIfFail {
					action = "delete"
				}

			}
			if out != "" {
				j.LastOutput = out
			}
			if j.DeleteIfSuccess {
				action = "delete"
				continue
			}
			return j.Save(h.storage)
		case "delete":
			//TODO: LOG DELETE
			fmt.Println("DELETE", j.ID)
			if err := j.Delete(); err != nil {
				return err
			}
			return nil
		case "skip":
			//TODO: LOG SKIP
			//fmt.Println("SKIP", j.ID)
			return nil
		case "done":
			//fmt.Println("DONE", j.ID)
			return nil
		}
	}
	return fmt.Errorf("unexpected out")
}

func Start(inStorage string, opts ...Option) error {
	cycle := 5
	found := 0
	jobErr := 0
	for _, o := range opts {
		if o.cycle > 0 {
			cycle = o.cycle
		}
	}
	h := New(inStorage).With(opts...)
	for {
		found = 0
		jobErr = 0
		fi, err := os.ReadDir(inStorage)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		for _, f := range fi {
			if f.IsDir() {
				continue
			}
			jb, err := job.Load(inStorage + f.Name())
			if err != nil {
				fmt.Println(f.Name(), err.Error())
				continue
			}
			found++
			if err := h.Handle(jb); err != nil {
				jobErr++
				fmt.Println(err.Error())
			}
		}
		fmt.Printf("jobs: %v (%v)  sleep for %v seconds...\r", found, jobErr, cycle)
		time.Sleep(time.Second * time.Duration(cycle))
	}
	return nil
}
