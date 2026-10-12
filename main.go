package main

import (
	"fmt"

	"task-tracker/models"
	"task-tracker/storage"
)

func main() {
	var store storage.Storage = storage.NewMemoryStorage()

	task, err := store.CreateTask(models.Task{
		Title: "blabla",
		Description: "bleble",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Создана: ", task)

	
	found, err := store.GetTask(task.ID)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Найдена: ", found)

	
	found.Title = "blublu"
	found.Status = "cooking"

	if err := store.UpdateTask(found); err != nil {
		fmt.Println(err)
		return
	}

	updated, _ := store.GetTask(task.ID)
	fmt.Println("Обновлена: ", updated)

	if err := store.DeleteTask(task.ID); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Задача удалена")
}