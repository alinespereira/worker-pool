package main

import (
	"math/rand"
	"time"

	"github.com/alinespereira/worker-pool/pool"
	"github.com/alinespereira/worker-pool/tasks"
)

func main() {
	wp := pool.NewWorkerPool(5)
	wp.Run()

	for range 10 {
		d := time.Duration(rand.Intn(10_000)) * time.Millisecond
		t := tasks.NewWaitTask(d)
		wp.Submit(t)
	}

	wp.Stop()
}
