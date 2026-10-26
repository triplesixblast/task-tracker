package storage

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"task-tracker/internal/dto"
)

type PostgresStorage struct {
	pool *pgxpool.Pool
}

func NewPostgresStorage(pool *pgxpool.Pool) *PostgresStorage {
	return &PostgresStorage{pool: pool}
}

func (s *PostgresStorage) CreateTask(ctx context.Context, task dto.Task) (dto.Task, error) {
	if task.Status == "" {
		task.Status = dto.StatusToDo
	}

	err := s.pool.QueryRow(ctx,
		`INSERT INTO tasks (title, description, due_date, status, priority)
		 VALUES ($1, $2, $3, $4, $5)
		 RETURNING id, created_at, updated_at`,
		task.Title, task.Description, task.DueDate, task.Status, task.Priority,
	).Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		return dto.Task{}, fmt.Errorf("create task: %w", err)
	}

	return task, nil
}

func (s *PostgresStorage) GetTask(ctx context.Context, id int64) (dto.Task, error) {
	var task dto.Task

	err := s.pool.QueryRow(ctx,
		`SELECT id, title, description, due_date, created_at, updated_at, status, priority
		 FROM tasks WHERE id = $1`,
		id,
	).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate,
		&task.CreatedAt, &task.UpdatedAt, &task.Status, &task.Priority)
	if errors.Is(err, pgx.ErrNoRows) {
		return dto.Task{}, ErrTaskNotFound
	}
	if err != nil {
		return dto.Task{}, fmt.Errorf("get task: %w", err)
	}

	return task, nil
}

func (s *PostgresStorage) UpdateTask(ctx context.Context, id int64, task dto.Task) error {
	if task.Status == "" {
		task.Status = dto.StatusToDo
	}

	tag, err := s.pool.Exec(ctx,
		`UPDATE tasks
		 SET title = $2, description = $3, due_date = $4, status = $5, priority = $6, updated_at = now()
		 WHERE id = $1`,
		id, task.Title, task.Description, task.DueDate, task.Status, task.Priority,
	)
	if err != nil {
		return fmt.Errorf("update task: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}

func (s *PostgresStorage) DeleteTask(ctx context.Context, id int64) error {
	tag, err := s.pool.Exec(ctx, `DELETE FROM tasks WHERE id = $1`, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}
	if tag.RowsAffected() == 0 {
		return ErrTaskNotFound
	}

	return nil
}
