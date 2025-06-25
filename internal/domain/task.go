package domain

import (
	"github.com/google/uuid"
	"time"
)

type TaskStatus string

const (
	StatusPending    TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted  TaskStatus = "completed"
	StatusFailed     TaskStatus = "failed"
)

type Task struct {
	ID          uuid.UUID  `json:"id"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	StartedAt   *time.Time `json:"started_at,omitempty"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
	Result      string     `json:"result,omitempty"`
	Error       string     `json:"error,omitempty"`
}

func NewTask() *Task {
	return &Task{
		ID:        uuid.New(),
		Status:    StatusPending,
		CreatedAt: time.Now(),
	}
}

func (t *Task) MarkAsProcessing() {
	t.Status = StatusProcessing
	now := time.Now()
	t.StartedAt = &now
}

func (t *Task) MarkAsCompleted(result string) {
	t.Status = StatusCompleted
	t.Result = result
	now := time.Now()
	t.CompletedAt = &now
}

func (t *Task) MarkAsFailed(err string) {
	t.Status = StatusFailed
	t.Error = err
	now := time.Now()
	t.CompletedAt = &now
}

func (t *Task) IsCompleted() bool {
	return t.Status == StatusCompleted || t.Status == StatusFailed
}
