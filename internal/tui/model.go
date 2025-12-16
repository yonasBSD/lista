package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/lista/internal/models"
)

type model struct {
	todoList *models.TodoList
	cursor   int // which todo is selected
	width    int // terminal width
	height   int // terminal height
}

func NewModel(todoList *models.TodoList) model {
	return model{
		todoList: todoList,
		cursor:   0,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}
