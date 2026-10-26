package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"

	"task-tracker/internal/app"
	"task-tracker/internal/storage"
)

const defaultDatabaseURL = "postgres://postgres:postgres@localhost:5432/task_tracker?sslmode=disable"

func main() {
	ctx := context.Background()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = defaultDatabaseURL
	}

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(fmt.Errorf("cant create db pool: %w", err))
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(fmt.Errorf("cant connect to db: %w", err))
	}

	store := storage.NewPostgresStorage(pool)

	app := app.BuildApp(store)

	if err := app.Start(":8080"); err != nil {
		panic(fmt.Errorf("cant start server: %w", err))
	}
}
