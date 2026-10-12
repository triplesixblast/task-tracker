package storage

import "task-tracker/models"

type Storage interface {
	CreateTask(task models.Task) (models.Task, error)
	GetTask(id int) (models.Task, error)
	UpdateTask(task models.Task) error
	DeleteTask(id int) error
}