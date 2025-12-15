package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var viewCmd = &cobra.Command{
	Use:   "view [id]",
	Short: "view a todo",
	Long:  "view a todo in its entirety, with notes if present",
	Args:  cobra.MinimumNArgs(1),
	Run:   viewTodo,
}

func viewTodo(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting id to int: %v", err)
		return
	}

	todos := todoList.List()

	for _, todo := range todos {
		if id == todo.ID {
			fmt.Printf("ID: %v\n", todo.ID)
			fmt.Printf("Title: %v\n", todo.Title)
			fmt.Printf("Priority: %v\n", todo.Priority)
			fmt.Printf("Status: %v\n", todo.Completed)
			fmt.Printf("Notes: \n%v\n", todo.Notes)
			return
		}
	}

	fmt.Printf("No todo with ID %v exists\n", id)
}
