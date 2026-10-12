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
	log.Printf("[Worker #%02d] Started", w.id)

	defer func() {
		wg.Done()
		log.Printf("[Worker #%02d] Done", w.id)
	}()

	for task := range taskCh {
		log.Printf("[Worker #%02d] Starting task", w.id)
		task.Run()
		log.Printf("[Worker #%02d] Task done", w.id)
	}
}
