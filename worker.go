package taskqueue

import (
	"context"
	"log/slog"
	"time"
)

type Worker struct {
	ID       int
	queue    *Queue
	handler  Handler
	retries  int
	interval time.Duration
}

func NewWorker(id int, q *Queue, h Handler, retries int) *Worker {
	return &Worker{
		ID:       id,
		queue:    q,
		handler:  h,
		retries:  retries,
		interval: time.Second,
	}
}

func (w *Worker) Start(ctx context.Context) {
	slog.Info("worker started", "id", w.ID)
	for {
		select {
		case <-ctx.Done():
			slog.Info("worker stopping", "id", w.ID)
			return
		default:
			task, err := w.queue.Dequeue(ctx)
			if err != nil {
				time.Sleep(w.interval)
				continue
			}
			if err := w.handler.Process(ctx, task); err != nil {
				slog.Error("task failed", "worker", w.ID, "task", task.ID, "err", err)
				w.queue.Retry(ctx, task, w.retries)
			}
		}
	}
}
