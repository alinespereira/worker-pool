package tasks

import (
	"context"
	"log"
	"time"
	"uuid"
)

type WaitTask struct {
	ID       uuid.UUID
	Duration time.Duration
}

func NewWaitTask(time time.Duration) WaitTask {
	return WaitTask{
		ID:       uuid.NewV4(),
		Duration: time,
	}
}

func (t WaitTask) Run(ctx context.Context) error {
	log.Printf("[Task %s] Starting...\n", t.ID)

	timer := time.NewTimer(t.Duration)
	defer timer.Stop()

	select {
	case <-timer.C:
		log.Printf("[Task %s] Done after %s\n", t.ID, t.Duration)
		return nil
	case <-ctx.Done():
		log.Printf("[Task %s] Cancelled: %v\n", t.ID, ctx.Err())
		return ctx.Err()
	}
}
