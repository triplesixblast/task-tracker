package storage

import (
	"context"
	"errors"
	"time"

	"task-tracker/internal/dto"
)

var ErrTaskNotFound = errors.New("task not found")

type MemoryStorage struct {
	taskStorage map[int64]dto.Task
	nextID      int64
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		taskStorage: make(map[int64]dto.Task),
		nextID:      1,
	}
}

func (s *MemoryStorage) CreateTask(_ context.Context, task dto.Task) (dto.Task, error) {
	task.ID = s.nextID
	s.nextID++

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if task.Status == "" {
		task.Status = dto.StatusToDo
	}

	s.taskStorage[task.ID] = task
	return task, nil
}

func (s *MemoryStorage) GetTask(_ context.Context, id int64) (dto.Task, error) {
	task, exists := s.taskStorage[id]
	if !exists {
		return dto.Task{}, ErrTaskNotFound
	}

	return task, nil
}

func (s *MemoryStorage) UpdateTask(_ context.Context, id int64, task dto.Task) error {
	oldTask, exists := s.taskStorage[id]
	if !exists {
		return ErrTaskNotFound
	}

	task.ID = id
	task.CreatedAt = oldTask.CreatedAt
	task.UpdatedAt = time.Now()

	s.taskStorage[id] = task

	return nil
}

func (s *MemoryStorage) DeleteTask(_ context.Context, id int64) error {
	if _, exists := s.taskStorage[id]; !exists {
		return ErrTaskNotFound
	}

	delete(s.taskStorage, id)

	return nil
}
