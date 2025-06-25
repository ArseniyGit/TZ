package repository

import (
	"TZ-API/internal/domain"
	"context"
	"github.com/google/uuid"
	"sync"
)

type MemoryTaskRepository struct {
	tasks map[uuid.UUID]*domain.Task
	mutex sync.RWMutex
}

func NewMemoryTaskRepository() *MemoryTaskRepository {
	return &MemoryTaskRepository{
		tasks: make(map[uuid.UUID]*domain.Task),
	}
}

func (r *MemoryTaskRepository) Create(ctx context.Context, task *domain.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.tasks[task.ID] = task
	return nil
}

func (r *MemoryTaskRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	task, exists := r.tasks[id]
	if !exists {
		return nil, nil
	}

	return task, nil
}

func (r *MemoryTaskRepository) Update(ctx context.Context, task *domain.Task) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[task.ID]; !exists {
		return domain.ErrTaskNotFound
	}

	r.tasks[task.ID] = task
	return nil
}

func (r *MemoryTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if _, exists := r.tasks[id]; !exists {
		return domain.ErrTaskNotFound
	}

	delete(r.tasks, id)
	return nil
}

func (r *MemoryTaskRepository) List(ctx context.Context) ([]*domain.Task, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	tasks := make([]*domain.Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}
