package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list tasks",
	Long:  "List all todo tasks in the todo list",
	Args:  cobra.MinimumNArgs(0),
	Run:   listTodos,
}

func listTodos(cmd *cobra.Command, args []string) {
	todos := todoList.List()

	fmt.Printf("%-4s %-10s %-50s\n", "ID", "STATUS", "DESCRIPTION")
	fmt.Println(strings.Repeat("-", 90))
	for _, todo := range todos {
		status := "PENDING"
		if todo.Completed {
			status = "DONE"
		}
		fmt.Printf("%-4d %-10s %-50s\n",
			todo.ID,
			status,
			todo.Text)
	}
}
