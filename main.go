package main

import (
	"fmt"
	"github.com/kwame-Owusu/todo-cli/internal/models"
	"github.com/kwame-Owusu/todo-cli/internal/storage"
)

func main() {
	todos := models.NewTodoList()
	todos.Add("This is first Todo")
	err := storage.SaveTodos(todos, "test.json")
	if err != nil {
		fmt.Printf("Error saving todos in main, %s", err)
	}

	loadedTodos, err := storage.LoadTodos("test.json")
	if err != nil {
		fmt.Printf("Error loading todos in main, %s", err)
	}

	fmt.Println(loadedTodos[0].Text)
}
