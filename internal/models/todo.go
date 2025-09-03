package models

import (
	"fmt"
)

type Todo struct {
	ID        int
	Text      string
	Completed bool
}

type TodoList struct {
	todos  []Todo
	nextID int
}

func NewTodoList() *TodoList {
	return &TodoList{
		todos:  make([]Todo, 0),
		nextID: 1,
	}
}

func (tl *TodoList) Add(text string) {
	todo := Todo{
		ID:        tl.nextID,
		Text:      text,
		Completed: false,
	}
	tl.todos = append(tl.todos, todo)
	tl.nextID++
}

func (tl *TodoList) Complete(id int) error {
	for i, todo := range tl.todos {
		if todo.ID == id {
			tl.todos[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}

func (tl *TodoList) List() []Todo {
	return tl.todos
}

func (tl *TodoList) Delete(id int) error {
	for i, todo := range tl.todos {
		if todo.ID == id {
			tl.todos = append(tl.todos[:i], tl.todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("todo with ID %d not found", id)
}
