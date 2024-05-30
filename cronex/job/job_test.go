package job

import (
	"fmt"
	"testing"

	"github.com/Galdoba/devtools/cronex/job/schedule"
)

func TestJob(t *testing.T) {
	job := Create("ffmpeg", "-i", `C:\Users\Admin\go\src\github.com\Galdoba\hello\go.mod`).
		SetSchedule("60/6 9-18 * * *")

	schd, err := schedule.New(job.ScheduleStr)
	if err != nil {
		t.Errorf(err.Error())
	}
	job.Schedule = schd
	fmt.Println(job.Schedule.Allowance(schedule.DayOfMonth))

	if err := job.Save(`C:\Users\Admin\go\src\github.com\Galdoba\devtools\cronex`); err != nil {
		t.Errorf("generation error: %v: %v", job.ID, err.Error())
	}
}
