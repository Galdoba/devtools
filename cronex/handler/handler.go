package handler

import "github.com/Galdoba/devtools/cronex/job"

type handler struct {
	logger   Logger
	Storage  string
	JobsByID map[string]*job.Job
}

type Logger interface {
}

func New(storage string) *handler {
	h := handler{}
	h.Storage = storage
	return &h
}

func (h *handler) Save(j *job.Job) error {
	return j.Save()
}

func (h *handler) Load(id string) (*job.Job, error) {
	path := h.Storage + id + ".json"
	return job.Load(path)
}

/*
Save(job)
Load(id) job

*/
