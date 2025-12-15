package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"sort"
	"strings"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list todos",
	Long:  "List all todo tasks in the todo list",
	Args:  cobra.MinimumNArgs(0),
	Run:   listTodos,
}

func listTodos(cmd *cobra.Command, args []string) {
	todos := todoList.List()

	// Sort: completed tasks at bottom, then by priority (high to low)
	sort.Slice(todos, func(i, j int) bool {
		// If one is completed and other isn't, completed goes to bottom
		if todos[i].Completed != todos[j].Completed {
			return !todos[i].Completed // false (not completed) comes first
		}
		// Both have same completion status, sort by priority (descending)
		return todos[i].Priority > todos[j].Priority
	})

	fmt.Printf("%-4s %-10s %-10s %-50s\n", "ID", "STATUS", "PRIORITY", "TITLE")
	fmt.Println(strings.Repeat("-", 70))

	for _, todo := range todos {
		status := "PENDING"
		title := todo.Title
		if todo.Completed {
			status = "DONE"
		}
		if len(todo.Notes) != 0 {
			title = fmt.Sprintf("%v 🗒️", title)
		}
		fmt.Printf("%-4d %-10s %-10s %-50s\n",
			todo.ID,
			status,
			todo.Priority,
			title)
	}
}
