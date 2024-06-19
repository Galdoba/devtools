package manager

import (
	"fmt"

	"github.com/Galdoba/devtools/app/hive/task"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("I select worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("I update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("I send work to worker")
}
