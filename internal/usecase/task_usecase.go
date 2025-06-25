package usecase

import (
	"TZ-API/internal/domain"
	"context"
	"github.com/google/uuid"
	"log"
)

type TaskUseCase struct {
	taskRepo      domain.TaskRepository
	taskProcessor domain.TaskProcessor
}

func NewTaskUseCase(taskRepo domain.TaskRepository, taskProcessor domain.TaskProcessor) *TaskUseCase {
	return &TaskUseCase{
		taskRepo:      taskRepo,
		taskProcessor: taskProcessor,
	}
}

func (uc *TaskUseCase) CreateTask(ctx context.Context) (*domain.Task, error) {
	task := domain.NewTask()

	if err := uc.taskRepo.Create(ctx, task); err != nil {
		return nil, err
	}

	go func() {
		if err := uc.taskProcessor.ProcessTask(context.Background(), task.ID); err != nil {
			log.Printf("Failed to process task %s: %v", task.ID, err)
		}
	}()

	return task, nil
}

func (uc *TaskUseCase) GetTask(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	task, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, domain.TaskNotFoundError(id)
	}

	return task, nil
}

func (uc *TaskUseCase) DeleteTask(ctx context.Context, id uuid.UUID) error {
	task, err := uc.taskRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if task == nil {
		return domain.TaskNotFoundError(id)
	}

	return uc.taskRepo.Delete(ctx, id)
}

func (uc *TaskUseCase) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return uc.taskRepo.List(ctx)
}
