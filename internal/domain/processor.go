package domain

import (
	"context"
	"github.com/google/uuid"
)

type TaskProcessor interface {
	ProcessTask(ctx context.Context, taskID uuid.UUID) error
}
