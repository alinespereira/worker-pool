package pool

import (
	"log"
	"sync"

	"github.com/alinespereira/worker-pool/tasks"
)

type worker struct {
	id uint
}

func newWorker(id uint) *worker {
	return &worker{
		id: id,
	}
}

func (w *worker) Run(wg *sync.WaitGroup, taskCh chan tasks.Task) {
	defer wg.Done()
	for task := range taskCh {
		log.Printf("[Worker #%2d] Starting task", w.id)
		task.Run()
		log.Printf("[Worker #%2d] Task done", w.id)
	}
}
