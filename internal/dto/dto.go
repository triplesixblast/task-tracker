package dto

import "time"

const (
	StatusToDo Status = "TODO"
	StatusWIP  Status = "WIP"
	StatusDone Status = "DONE"
)

type Status = string

type Task struct {
	ID          int64
	Title       string
	Description string
	DueDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      Status
	Priority    int
}
