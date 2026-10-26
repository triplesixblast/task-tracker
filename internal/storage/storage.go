package storage

import (
	"context"
	"errors"

	"task-tracker/internal/dto"
)

var ErrTaskNotFound = errors.New("task not found")

type Storage interface {
	CreateTask(context.Context, dto.Task) (dto.Task, error)
	GetTask(context.Context, int64) (dto.Task, error)
	UpdateTask(context.Context, int64, dto.Task) error
	DeleteTask(context.Context, int64) error
}
