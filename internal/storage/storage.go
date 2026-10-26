package storage

import (
	"context"

	"task-tracker/internal/dto"
)

type Storage interface {
	CreateTask(context.Context, dto.Task) (dto.Task, error)
	GetTask(context.Context, int64) (dto.Task, error)
	UpdateTask(context.Context, int64, dto.Task) error
	DeleteTask(context.Context, int64) error
}
