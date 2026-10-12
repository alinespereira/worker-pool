package tasks

import (
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

func (t WaitTask) Run() {
	log.Printf("[Task %s] Starting...\n", t.ID)
	time.Sleep(t.Duration)
	log.Printf("[Task %s] Done after %s\n", t.ID, t.Duration)
}
