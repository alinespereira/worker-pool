package pool

import (
	"context"
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

func (w *worker) Run(ctx context.Context, wg *sync.WaitGroup, taskCh chan tasks.Task) {
	log.Printf("[Worker #%02d] Started", w.id)

	defer func() {
		wg.Done()
		log.Printf("[Worker #%02d] Done", w.id)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Printf("[Worker #%02d] Stopping: %v", w.id, ctx.Err())
			return
		case task, ok := <-taskCh:
			if !ok {
				return
			}
			log.Printf("[Worker #%02d] Starting task", w.id)
			if err := task.Run(ctx); err != nil {
				log.Printf("[Worker #%02d] Task failed: %v", w.id, err)
				continue
			}
			log.Printf("[Worker #%02d] Task done", w.id)
		}
	}
}
