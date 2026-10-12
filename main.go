package main

import (
	"context"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alinespereira/worker-pool/pool"
	"github.com/alinespereira/worker-pool/tasks"
)

func main() {
	// Cancel on Ctrl+C / SIGTERM, or after 5 seconds, whichever comes first.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	wp := pool.NewWorkerPool(5)
	wp.Run(ctx)

	for range 10 {
		d := time.Duration(rand.Intn(10_000)) * time.Millisecond
		t := tasks.NewWaitTask(d)
		if err := wp.Submit(ctx, t); err != nil {
			log.Printf("Could not submit task: %v", err)
			break
		}
	}

	wp.Stop()
}
