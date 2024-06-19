package worker

import (
	"fmt"

	"github.com/Galdoba/devtools/app/hive/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("I collect stats")
}

func (w *Worker) RunTasks() {
	fmt.Println("I start or stop tasks")
}

func (w *Worker) StartTask() {
	fmt.Println("I start task")
}

func (w *Worker) StopTask() {
	fmt.Println("I stop task")
}
