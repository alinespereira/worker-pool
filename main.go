package main

import (
	"math/rand"
	"time"

	"github.com/alinespereira/worker-pool/pool"
	"github.com/alinespereira/worker-pool/tasks"
)

func main() {
	wp := pool.NewWorkerPool(5)

	for range 10 {
		d := time.Duration(rand.Intn(10_000)) * time.Second
		t := tasks.NewWaitTask(d)
		wp.Submit(t)
	}

	wp.Stop()
}
