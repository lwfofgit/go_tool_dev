package work_pool

import "code.byted.org/gopkg/pkg/log"

type WorkPool struct {
	WorkerCapacity int
	Job            chan *Task
	Result         chan *Result
	Quit           chan bool
}

func NewWorkPool(workerCapacity int) *WorkPool {
	wp := WorkPool{
		WorkerCapacity: workerCapacity,
		Job:            make(chan *Task),
		Result:         make(chan *Result),
	}
	return &wp
}

func (wp *WorkPool) Run() {
	if wp.WorkerCapacity == 0 {
		return
	}
	for i := 0; i < wp.WorkerCapacity; i++ {
		log.Infoln("........worker %d", i)
		go wp.Worker()
	}
}

func (wp *WorkPool) Worker() {
	log.Infoln("........worker")
	for {
		select {
		case job := <-wp.Job:
			if job.handler == nil {
				log.Infoln("........return")
				return
			}
			err := job.handler()
			if err != nil {
				log.Errorf("job execute failed, err: %s", err)
			}
		case <-wp.Quit:
			{
				return
			}

		}
	}
}

func (wp *WorkPool) Stop() {
	go func() {
		wp.Quit <- true
	}()
}

func (wp *WorkPool) Submit(task *Task) {
	if task == nil {
		return
	}
	wp.Job <- task
}
