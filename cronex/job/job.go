package job

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Galdoba/devtools/cronex/job/schedule"
)

/*
# ┌────────────── second (0-59)
# | ┌───────────── minute (0–59)
# | │ ┌───────────── hour (0–23)
# | │ │ ┌───────────── day of the month (1–31)
# | │ │ │ ┌───────────── month (1–12)
# | │ │ │ │ ┌───────────── day of the week (0–6) (Sunday to Saturday;
# | │ │ │ │ │                                   7 is also Sunday on some systems)
# | │ │ │ │ │
# | │ │ │ │ │
# * * * * * * <command to execute>
*/

type Job struct {
	ID               string   `json:"ID,omitempty"`                //timestamp time.Now().UnixNano()
	ScheduleStr      string   `json:"Schedule"`                    //* * * * *
	Repeatable       bool     `json:"Repeatable,omitempty"`        //repeat after first handling
	DeleteImmediatly bool     `json:"Delete Immediatly,omitempty"` //delete upon touch
	InProgress       bool     `json:"In Progress,omitempty"`       //flag for handler
	Skip             bool     `json:"Skip,omitempty"`              //flag for handler: do not start atempt but log as event
	Halt             bool     `json:"Halt,omitempty"`              //flag for handler: do not start atempt, do not log as event (overwrite SKIP)
	Atempts          int      `json:"Total Atempts,omitempty"`     //Atempt Counter
	Done             int      `json:"Total Successes,omitempty"`   //Success Counter
	StartDir         string   `json:"Run From,omitempty"`
	Handler          string   `json:"Handler utility"`
	Args             []string `json:"Args,omitempty"`
	LastOutput       string   `json:"Last output,omitempty"`
	// second           []int
	// minute           []int
	// hour             []int
	// doM              []int
	// month            []int
	// doW              []int
	ready    bool
	saveTo   string //куда сохранить json
	errMsg   string
	Schedule schedule.Schedule
}

// Newjob - создает объект Job
func Create(storage string, agent string, args ...string) *Job {
	job := Job{}
	job.ID = newJobID()
	job.Handler = agent
	job.Args = args
	job.errMsg = "no validation commenced"
	job.saveTo = storage + string(filepath.Separator) + job.ID + ".json"
	return &job
}

// SetSchedule - НЕОБХОДИМАЯ ФУНКЦИЯ: задает параметры расписания для хендлера
func (job *Job) SetSchedule(schdStr string) *Job {
	job.ScheduleStr = schdStr
	err := fmt.Errorf("schedule not injected")
	sched, err := schedule.New(job.ScheduleStr)
	if err != nil {
		job.errMsg = err.Error()
		return job
	}
	job.Schedule = sched
	return job
}

func (job *Job) Doer() string {
	return job.Handler
}

// newJobID - задает уникальный ID для Job
func newJobID() string {
	str := time.Now().Format("20060102150405.999")
	data := strings.Split(str, ".")
	if len(data) < 2 {
		data = append(data, "")
	}
	for len(data[1]) < 3 {
		data[1] += "0"
	}
	return "cronex_job_" + strings.Join(data, "")
}

// validate - проверяет мелкие недочеты и ставит маркер готовности
func (job *Job) validate() {
	if job.errMsg == "no validation commenced" {
		job.errMsg = ""
	}
	if job.errMsg != "" {
		return
	}
	if !pathIsValid(job.saveTo) {
		job.errMsg = "path invalid: " + job.saveTo
		return
	}
	if job.ScheduleStr == "" {
		job.errMsg = "job has no schedule"
		return
	}
	job.ready = true
}

func (job *Job) convertSchedule() error {
	return nil
}

func writeAsJson(job *Job) error {
	fileName := job.saveTo + string(filepath.Separator) + job.ID + ".json"
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	defer f.Close()
	if err != nil {
		return fmt.Errorf("open file: %v", err.Error())
	}
	bt, err := json.MarshalIndent(job, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal job: %v", err.Error())
	}
	f.Truncate(0)
	_, err = f.Write(bt)
	if err != nil {
		return fmt.Errorf("write to file: %v", err.Error())
	}
	return nil
}

func pathIsValid(path string) bool {
	dir := filepath.Dir(path)
	_, err := os.ReadDir(dir)
	if err != nil {
		return false
	}
	return true
}
