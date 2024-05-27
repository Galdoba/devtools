package cronex

import (
	"testing"
)

func TestJob(t *testing.T) {
	job := NewJob("echo", "hello world!").
		SetSchedule("60/4,9,10,13-26 * * * *")

	if err := job.Save(); err != nil {
		t.Errorf("generation error: %v: %v", job.ID, err.Error())
	}
}
