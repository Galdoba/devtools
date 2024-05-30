package job

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Galdoba/devtools/cli/command"
)

func (job *Job) Save(path string) error {
	job.saveTo = path
	job.validate()
	if !job.ready {
		return fmt.Errorf(job.errMsg)
	}
	//check directory
	if !pathIsValid(job.saveTo) {
		return fmt.Errorf("save path is not valid: %v", job.saveTo)
	}
	if err := writeAsJson(job); err != nil {
		return fmt.Errorf("write to file: %v", err.Error())
	}
	return nil
}

func (job *Job) Delete() error {
	fmt.Printf("%v\n", job.saveTo+string(filepath.Separator)+job.ID+".json")
	return os.Remove(job.saveTo + string(filepath.Separator) + job.ID + ".json")
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

func (j *Job) Execute() (string, error) {

	j.Atempts++
	fmt.Println("run:", j.Handler, j.Args)
	cline := j.Handler + " " + strings.Join(j.Args, " ")
	out, er, err := command.Execute(cline, command.Set(command.TERMINAL_ON), command.Set(command.BUFFER_ON))
	//TODO: LOG EXECUTION
	if err != nil {
		return er, err
	}
	j.Done++
	return out, nil
}
