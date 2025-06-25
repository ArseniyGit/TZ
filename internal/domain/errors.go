package domain

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrTaskNotFound          = errors.New("task not found")
	ErrTaskAlreadyProcessing = errors.New("task is already being processed")
)

func TaskNotFoundError(id uuid.UUID) error {
	return fmt.Errorf("task with id %s not found", id.String())
}
