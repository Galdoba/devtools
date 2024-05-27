package cronex

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
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
	ID              string   `json:"ID,omitempty"`                      //timestamp time.Now().UnixNano()
	Schedule        string   `json:"Schedule"`                          //timestamp time.Now().UnixNano()
	DoOnce          bool     `json:"Do Once,omitempty"`                 //no repeat after success
	DeleteIfSuccess bool     `json:"Delete after completion,omitempty"` //delete after success
	DeleteIfAtempt  bool     `json:"Delete after atempt,omitempty"`     //delete after atempt
	InProgress      bool     `json:"In Progress,omitempty"`             //flag for handler
	Skip            bool     `json:"Skip,omitempty"`                    //flag for handler: do not start atempt but log as event
	Halt            bool     `json:"Halt,omitempty"`                    //flag for handler: do not start atempt, do not log as event (overwrite SKIP)
	Atempts         int      `json:"Total Atempts,omitempty"`           //Atempt Counter
	Done            int      `json:"Total Successes,omitempty"`         //Success Counter
	StartDir        string   `json:"Run From,omitempty"`
	Handler         string   `json:"Handler utility"`
	Args            []string `json:"Args,omitempty"`

	minute []int
	hour   []int
	doM    []int
	month  []int
	doW    []int
	ready  bool
	saveTo string //куда сохранить json
	errMsg string
}

//Newjob - создает объект Job
func NewJob(handler string, args ...string) *Job {
	job := Job{}
	job.ID = newJobID()
	job.Handler = handler
	job.Args = args
	job.errMsg = "no validation commenced"
	return &job
}

//SetSchedule - НЕОБХАДИМАЯ ФУНКЦИЯ: задает параметры расписания для хендлера
func (job *Job) SetSchedule(schedule string) *Job {
	job.Schedule = schedule
	err := fmt.Errorf("schedule not injected")
	job.minute, job.hour, job.doM, job.month, job.doW, err = scheduleToData(job.Schedule)
	if err != nil {
		job.errMsg = fmt.Sprintf("schedule: %v", err.Error())
	}
	return job
}

//newJobID - задает уникальный ID для Job
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

//validate - проверяет мелкие недочеты и ставит маркер готовности
func (job *Job) validate() {
	if job.errMsg == "no validation commenced" {
		job.errMsg = ""
	}
	if job.errMsg != "" {
		return
	}
	if job.saveTo == "" {
		here, err := os.Executable()
		path := filepath.Dir(here)
		if err != nil {
			job.errMsg = fmt.Sprintf("failed to get current directory: %v", err.Error())
			return
		}
		job.saveTo = path
	}
	if job.Schedule == "" {
		job.errMsg = "job has no schedule"
		return
	}
	job.ready = true
}

func (job *Job) convertSchedule() error {
	return nil
}

func (job *Job) Save() error {
	job.validate()
	if !job.ready {
		return fmt.Errorf(job.errMsg)
	}
	fmt.Println("check directory", job.saveTo)
	fmt.Println("Save Job as", job.saveTo+string(filepath.Separator)+job.ID+".json")
	return nil
}

func scheduleToData(sch string) ([]int, []int, []int, []int, []int, error) {
	mm, hh, dom, mn, dow := []int{}, []int{}, []int{}, []int{}, []int{}
	min, max := 0, 0
	err := fmt.Errorf("not parsed")
	parts := strings.Split(sch, " ")
	if len(parts) != 5 {
		return mm, hh, dom, mn, dow, fmt.Errorf("expect to have 5 parts (have %v)", len(parts))
	}
	for i, part := range parts {
		switch i {
		case 0:
			min = 0
			max = 59
			mm, err = parseBlock(part, min, max)
		case 1:
			min = 0
			max = 23
			hh, err = parseBlock(part, min, max)
		case 2:
			min = 1
			max = 31
			dom, err = parseBlock(part, min, max)
		case 3:
			min = 1
			max = 12
			mn, err = parseBlock(part, min, max)
		case 4:
			min = 0
			max = 7
			dow, err = parseBlock(part, min, max)
		}
		if err != nil {
			return mm, hh, dom, mn, dow, fmt.Errorf("block %v: %v", i, err.Error())
		}
	}

	return mm, hh, dom, mn, dow, nil
}

func parseBlock(s string, min, max int) ([]int, error) {
	if s == "*" {
		return allNumbers(min, max), nil
	}
	if s == "" {
		return []int{}, fmt.Errorf("no data in block")
	}
	mm := []int{}
	segm := strings.Split(s, ",")
	for _, seg := range segm {
		if strings.Contains(seg, "/") {
			sp, err := splitDevision(seg)
			if err != nil {
				return mm, fmt.Errorf("parsing failed: %v", err.Error())
			}
			mm = append(mm, sp...)
			continue
		}
		if strings.Contains(seg, "-") {
			sp, err := splitOrderly(seg)
			if err != nil {
				return mm, fmt.Errorf("parsing failed: %v", err.Error())
			}
			mm = append(mm, sp...)
			continue
		}
		if val, err := strconv.Atoi(seg); err == nil {
			mm = append(mm, val)
			continue
		}
		return mm, fmt.Errorf("parsing failed: %v", s)
	}
	minutes, err := checkVals(mm, min, max)
	if err != nil {
		return mm, fmt.Errorf("parsing failed: %v", err.Error())
	}
	return minutes, nil
}

func splitOrderly(s string) ([]int, error) {
	sp := []int{}
	data := strings.Split(s, "-")
	if len(data) != 2 {
		return sp, fmt.Errorf("'%v': expect to have 2 numbers", s)
	}
	min, err := strconv.Atoi(data[0])
	if err != nil {
		return sp, fmt.Errorf("%v: first number: %v", s, err.Error())
	}
	max, err2 := strconv.Atoi(data[1])
	if err2 != nil {
		return sp, fmt.Errorf("%v: second number: %v", s, err2.Error())
	}
	for i := min; i <= max; i++ {
		sp = append(sp, i)
	}
	return sp, nil
}

func splitDevision(s string) ([]int, error) {
	sp := []int{}
	data := strings.Split(s, "/")
	if len(data) != 2 {
		return sp, fmt.Errorf("'%v': expect to have 2 numbers", s)
	}
	dev, err := strconv.Atoi(data[0])
	if err != nil {
		return sp, fmt.Errorf("'%v': first number: %v", s, err.Error())
	}
	dvtr, err2 := strconv.Atoi(data[1])
	if err2 != nil || dvtr == 0 {
		return sp, fmt.Errorf("'%v': can't devide by zero", s)
	}
	for i := 0; i < dev; i++ {
		if i%dvtr == 0 {
			sp = append(sp, i)
		}
	}
	return sp, nil
}

func allNumbers(min, max int) []int {
	sl := []int{}
	for i := min; i <= max; i++ {
		sl = append(sl, i)
	}
	return sl
}

func checkVals(sl []int, min, max int) ([]int, error) {
	formatted := []int{}
	if len(sl) < 1 {
		return formatted, fmt.Errorf("no values parsed")
	}
	sort.Ints(sl)
	if sl[0] < min {
		return formatted, fmt.Errorf("bad values parsed: %v", sl[0])
	}
	if sl[len(sl)-1] > max {
		return formatted, fmt.Errorf("bad values parsed: %v", sl[len(sl)-1])
	}
	for _, i := range sl {
		formatted = appendUnique(formatted, i)
	}
	return formatted, nil
}

func appendUnique(sl []int, elem int) []int {
	for _, i := range sl {
		if i == elem {
			return sl
		}
	}
	return append(sl, elem)
}
