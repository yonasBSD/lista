package main

import (
	"fmt"
	"github.com/kwame-Owusu/todo-cli/internal/models"
)

func main() {
	todos := models.NewTodoList()
	todos.Add("This is first Todo")
	fmt.Println(todos.List())
}
