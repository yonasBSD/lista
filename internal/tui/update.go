package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/kwame-Owusu/lista/internal/storage"
)

// msgTodoSaved is sent when saving finishes (success or error)
type msgTodoSaved struct {
	err error
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		// Confirm deletion
		case "y", "enter":
			cmd = m.deleteTodo()

		// Cancel deletion
		case "n", "esc":
			m.cancelDelete()

		// Quit app
		case "q", "ctrl+c":
			return m, tea.Quit

		// Cursor navigation
		case "up", "k":
			m.moveCursorUp()
		case "down", "j":
			m.moveCursorDown()

		// Toggle completion
		case " ":
			cmd = m.toggleTodo()

		// Prepare deletion confirmation
		case "d", "x":
			m.confirmDeleteAtCursor()
		}

	case msgTodoSaved:
		// Handle async save completion
		if msg.err != nil {
			m.err = msg.err
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	return m, cmd
}

// Helper methods

// deleteTodo deletes the selected todo and returns a Cmd to save asynchronously
func (m *model) deleteTodo() tea.Cmd {
	err := m.todoList.Delete(m.deleteID)
	if err != nil {
		m.err = err
	}

	m.confirmDelete = false
	m.deleteID = 0

	// Clamp cursor
	if m.cursor >= len(m.todoList.List()) && m.cursor > 0 {
		m.cursor--
	}

	// Return async save command
	return m.saveTodosCmd()
}

func (m *model) toggleTodo() tea.Cmd {
	todos := m.todoList.List()
	if len(todos) == 0 || m.cursor >= len(todos) {
		return nil
	}

	selectedID := todos[m.cursor].ID

	if err := m.todoList.Toggle(selectedID); err != nil {
		m.err = err
		return nil
	}

	// Re-sync cursor after any potential reordering
	todos = m.todoList.List()
	if idx := findTodoIndexByID(todos, selectedID); idx >= 0 {
		m.cursor = idx
	}

	return m.saveTodosCmd()
}

// saveTodosCmd returns a command that saves todos asynchronously
func (m *model) saveTodosCmd() tea.Cmd {
	todos := m.todoList.Todos
	filename := m.filename

	return func() tea.Msg {
		err := storage.SaveTodos(todos, filename)
		return msgTodoSaved{err: err}
	}
}

func (m *model) cancelDelete() {
	m.confirmDelete = false
	m.deleteID = 0
}

func (m *model) confirmDeleteAtCursor() {
	todos := m.todoList.List()
	if len(todos) == 0 || m.cursor >= len(todos) {
		return
	}

	m.confirmDelete = true
	m.deleteID = todos[m.cursor].ID
}

func (m *model) moveCursorUp() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *model) moveCursorDown() {
	if m.cursor < len(m.todoList.Todos)-1 {
		m.cursor++
	}
}
