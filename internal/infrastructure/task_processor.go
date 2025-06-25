package infrastructure

import (
	"TZ-API/internal/domain"
	"context"
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type TaskProcessor struct {
	taskRepo domain.TaskRepository
}

func NewTaskProcessor(taskRepo domain.TaskRepository) *TaskProcessor {
	return &TaskProcessor{
		taskRepo: taskRepo,
	}
}

func (p *TaskProcessor) ProcessTask(ctx context.Context, taskID uuid.UUID) error {
	task, err := p.taskRepo.GetByID(ctx, taskID)
	if err != nil {
		return err
	}

	if task == nil {
		return domain.TaskNotFoundError(taskID)
	}

	if task.Status == domain.StatusProcessing {
		return domain.ErrTaskAlreadyProcessing
	}

	task.MarkAsProcessing()
	if err := p.taskRepo.Update(ctx, task); err != nil {
		return err
	}

	processingTime := time.Duration(3+rand.Intn(3)) * time.Minute

	select {
	case <-time.After(processingTime):
		result := fmt.Sprintf("Task completed successfully after %s", processingTime)
		task.MarkAsCompleted(result)

	case <-ctx.Done():
		task.MarkAsFailed("Task processing was cancelled")

	case <-time.After(processingTime + time.Second*10):
		if rand.Float32() < 0.1 {
			task.MarkAsFailed("Random processing error occurred")
		} else {
			result := fmt.Sprintf("Task completed successfully after %s", processingTime)
			task.MarkAsCompleted(result)
		}
	}

	return p.taskRepo.Update(ctx, task)
}
