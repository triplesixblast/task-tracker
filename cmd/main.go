package main

import (
	"fmt"

	"task-tracker/internal/app"
)

func main() {
	app := app.BuildApp()

	if err := app.Start(":8080"); err != nil {
		panic(fmt.Errorf("cant start server: %w", err))
	}
}
