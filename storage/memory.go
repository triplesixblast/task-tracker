package storage

import (
	"errors"
	"time"

	"task-tracker/models"
)

type MemoryStorage struct {
	tasks map[int]models.Task
	nextID int
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		tasks: make(map[int]models.Task),
		nextID: 1,
	}
}

func (s *MemoryStorage) CreateTask(task models.Task) (models.Task, error) {
	task.ID = s.nextID
	s.nextID++

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	if task.Status == "" {
		task.Status = "New"
	}

	s.tasks[task.ID] = task
	return task, nil
}

func (s *MemoryStorage) GetTask(id int) (models.Task, error) {
	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("Task not found")
	}

	return task, nil
}

func (s *MemoryStorage) UpdateTask(task models.Task) error {
	oldTask, exists := s.tasks[task.ID]
	if !exists {
		return errors.New("Task not found")
	}

	task.CreatedAt = oldTask.CreatedAt
	task.UpdatedAt = time.Now()

	s.tasks[task.ID] = task

	return nil
}

func (s *MemoryStorage) DeleteTask(id int) error {
	if _, exists := s.tasks[id]; !exists {
		return errors.New("Task not found")
	}

	delete(s.tasks, id)

	return nil
}