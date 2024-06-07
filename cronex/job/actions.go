package job

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Galdoba/devtools/cli/command"
	"github.com/Galdoba/devtools/cronex/job/schedule"
)

func (job *Job) Save() error {
	job.validate()
	if !job.ready {
		return fmt.Errorf(job.errMsg)
	}
	if err := writeAsJson(job); err != nil {
		return fmt.Errorf("write to file: %v", err.Error())
	}
	return nil
}

func (job *Job) Delete() error {
	return os.Remove(job.saveTo)
}

func Load(path string) (*Job, error) {
	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("not a json")
	}
	bt, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file failed: %v", err.Error())
	}
	j := &Job{}
	err = json.Unmarshal(bt, j)
	if err != nil {
		return nil, fmt.Errorf("unmarshal failed: %v", err.Error())
	}
	j.SetSchedule(j.ScheduleStr)
	j.saveTo = filepath.Dir(path)
	return j, nil
}

type Logger interface {
	Info()
	Error()
}

func (j *Job) Execute() error {
	if j.InProgress {
		return nil
	}
	j.Atempts++
	j.InProgress = true
	cline := j.Handler + " " + strings.Join(j.Args, " ")
	out, er, err := command.Execute(cline, command.Set(command.TERMINAL_ON), command.Set(command.BUFFER_ON))
	//TODO: LOG EXECUTION
	if err != nil {
		fmt.Println(out, er)
		return err
	}
	j.InProgress = false
	j.Done++
	return nil
}

func (j *Job) TimeValid(t time.Time) bool {
	return schedule.TimeValid(j.Schedule, t)
}
