package pool

import (
	"context"
	"errors"
	"sync"

	"github.com/alinespereira/worker-pool/tasks"
)

// ErrPoolStopped is returned by Submit when the pool's context was cancelled
// and the workers are no longer accepting tasks.
var ErrPoolStopped = errors.New("worker pool stopped")

type WorkerPool struct {
	taskCh    chan tasks.Task
	workers   []*worker
	pool_size uint
	wg        *sync.WaitGroup
	ctx       context.Context
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
		ctx:       context.Background(),
	}
}

// Submit blocks until a worker picks up the task. It returns ctx.Err() if ctx
// is cancelled first, or ErrPoolStopped if the pool's context is cancelled.
func (wp *WorkerPool) Submit(ctx context.Context, task tasks.Task) error {
	select {
	case wp.taskCh <- task:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	case <-wp.ctx.Done():
		return ErrPoolStopped
	}
}

// Run starts the workers. When ctx is cancelled, workers stop taking new
// tasks and running tasks receive the cancellation through their context.
func (wp *WorkerPool) Run(ctx context.Context) {
	wp.ctx = ctx
	for _, worker := range wp.workers {
		wp.wg.Add(1)
		go worker.Run(ctx, wp.wg, wp.taskCh)
	}
}

// Stop closes the task channel and waits for all workers to finish.
func (wp *WorkerPool) Stop() {
	close(wp.taskCh)
	wp.wg.Wait()
}
