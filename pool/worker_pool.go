package pool

import (
	"sync"

	"github.com/alinespereira/worker-pool/tasks"
)

type WorkerPool struct {
	taskCh    chan tasks.Task
	workers   []*worker
	pool_size uint
	wg        *sync.WaitGroup
}

func NewWorkerPool(pool_size uint) WorkerPool {
	workers := []*worker{}

	for i := range pool_size {
		workers = append(workers, newWorker(i+1))
	}

	return WorkerPool{
		taskCh:    make(chan tasks.Task),
		workers:   workers,
		pool_size: pool_size,
		wg:        &sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Submit(task tasks.Task) {
	wp.taskCh <- task
}

func (wp *WorkerPool) Run() {
	for _, worker := range wp.workers {
		wp.wg.Add(1)
		worker.Run(wp.wg, wp.taskCh)
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.taskCh)
	wp.wg.Wait()
}
