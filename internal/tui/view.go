package tui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func (m model) View() string {
	var b strings.Builder

	// Title with style
	title := titleStyle.Render(`
	╔════════════════╗
	║     LISTA      ║
	╚════════════════╝
    Your CLI todo manager
`)
	b.WriteString(title + "\n")

	if m.err != nil {
		errorMsg := errorStyle.
			Render(fmt.Sprintf("⚠ Error: %v", m.err))
		b.WriteString(errorMsg + "\n\n")
	}

	todos := m.todoList.List()

	if len(todos) == 0 {
		empty := itemStyle.Render("No todos yet. Add one to get started!")
		b.WriteString(empty + "\n")
	}

	for i, todo := range todos {
		// Cursor indicator
		cursor := "  "
		if m.cursor == i {
			cursor = cursorStyle.Render("▶ ")
		} else {
			cursor = "  "
		}

		// Checkbox
		checkbox := "○"
		if todo.Completed {
			checkbox = "✓"
		}

		noteIndicator := ""
		if len(todo.Notes) > 0 {
			noteIndicator = "󰈙"
		}

		content := fmt.Sprintf(
			"%s %s [%s] %s",
			checkbox,
			todo.Title,
			todo.Priority,
			noteIndicator,
		)

		// Priority badge
		priorityBadge := getPriorityStyle(todo.Priority.String()).
			Render(fmt.Sprintf("[%s]", todo.Priority))

		// Todo title
		todoTitle := todo.Title

		// Build the line
		var line string
		if m.cursor == i {
			// Selected item - highlight entire line
			line = cursor + selectedStyle.Render(content)
			if todo.Completed {
				line = cursor + completedSelectedStyle.Render(content)
			}
		} else if todo.Completed {
			// Completed item
			line = cursor + completedStyle.Render(content)
		} else {
			// Normal item
			line = fmt.Sprintf("%s %s %s %s", cursor, checkbox,
				itemStyle.Render(todoTitle), priorityBadge)
		}

		b.WriteString(line + "\n")
	}

	if len(todos) != 0 {
		taskCount := helpStyle.Render("\n0 of 3 complete")
		b.WriteString(taskCount)
	}

	// Help text
	help := helpStyle.Render("\n↑/↓: navigate • space: toggle • a: add • d: delete • e: edit • q: quit")
	b.WriteString(help)

	if m.confirmDelete {
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

		// Overlay modal on top of UI
		return lipgloss.Place(
			m.width,
			m.height,
			lipgloss.Center,
			lipgloss.Center,
			modal,
		)
	}

	return b.String()
}
