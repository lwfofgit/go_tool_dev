package work_pool

import "fmt"

type Task struct {
	handler func() error
}

func NewTask(handler func() error) *Task {
	task := &Task{
		handler: handler,
	}
	return task
}

func (t *Task) Execute() error {
	err := t.handler()
	if err != nil {
		return fmt.Errorf("[task] task execute handler failed, err: %s", err)
	}
	return nil
}
