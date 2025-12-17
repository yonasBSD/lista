package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/kwame-Owusu/lista/internal/models"
)

func (m model) View() string {
	if m.confirmDelete {
		return m.renderDeleteModal()
	}

	var b strings.Builder

	b.WriteString(m.renderTitle())
	b.WriteString(m.renderError())
	b.WriteString(m.renderTodos())
	b.WriteString(m.renderCompletedCount())
	b.WriteString(m.renderHelp())

	return b.String()
}

func (m model) renderTitle() string {
	title := titleStyle.Render(`
	╔════════════════╗
	║     LISTA      ║
	╚════════════════╝
	Your CLI todo manager
`)
	return title + "\n"
}

func (m model) renderError() string {
	if m.err == nil {
		return ""
	}
	return errorStyle.Render(fmt.Sprintf("⚠ Error: %v", m.err)) + "\n\n"
}

func (m model) renderTodos() string {
	todos := m.todoList.List()
	if len(todos) == 0 {
		return itemStyle.Render("No todos yet. Add one to get started!") + "\n"
	}

	var b strings.Builder
	for i, todo := range todos {
		b.WriteString(m.renderTodoLine(i, todo) + "\n")
	}
	return b.String()
}

func (m model) renderTodoLine(i int, todo models.Todo) string {
	// Cursor
	cursor := "  "
	if m.cursor == i {
		cursor = cursorStyle.Render("▶ ")
	}

	// Checkbox
	checkbox := "○"
	if todo.Completed {
		checkbox = "✓"
	}

	// Note indicator
	noteIndicator := ""
	if len(todo.Notes) > 0 {
		noteIndicator = "󰈙"
	}

	content := fmt.Sprintf("%s %s [%s] %s", checkbox, todo.Title, todo.Priority, noteIndicator)
	priorityBadge := getPriorityStyle(todo.Priority.String()).Render(fmt.Sprintf("[%s]", todo.Priority))
	todoTitle := todo.Title

	// Determine style
	if m.cursor == i {
		if todo.Completed {
			return cursor + completedSelectedStyle.Render(content)
		}
		return cursor + selectedStyle.Render(content)
	} else if todo.Completed {
		return cursor + completedStyle.Render(content)
	}
	return fmt.Sprintf("%s %s %s %s", cursor, checkbox, itemStyle.Render(todoTitle), priorityBadge)
}

func (m model) renderCompletedCount() string {
	todos := m.todoList.List()
	if len(todos) == 0 {
		return ""
	}

	completedCount := len(m.todoList.GetCompleted())
	countString := fmt.Sprintf("%v of %v complete", completedCount, len(todos))
	return helpStyle.Render(countString)
}

func (m model) renderHelp() string {
	return helpStyle.Render("\n↑/↓: navigate • space: toggle • a: add • d: delete • e: edit • q: quit")
}

func (m model) renderDeleteModal() string {
	todos := m.todoList.List()
	var title string
	for _, t := range todos {
		if t.ID == m.deleteID {
			title = t.Title
			break
		}
	}

	modal := lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		modalStyle.Render(
			fmt.Sprintf(
				"Delete \"%s\"?\n\n%s",
				title,
				cursorStyle.Render("y: confirm • n / esc: cancel"),
			),
		),
	)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		modal,
	)
}
