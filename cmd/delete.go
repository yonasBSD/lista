package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete a Todo",
	Long:  "Delete a Todo and remove it from the list",
	Args:  cobra.MinimumNArgs(1),
	Run:   deleteTodo,
}

func deleteTodo(cmd *cobra.Command, args []string) {
	todoId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting argument into int\n")
		return
	}
	defer saveTodos()
	err = todoList.Delete(todoId)
	if err != nil {
		fmt.Printf("Error deleting todo with id: %s\n", err)
		return
	}
	fmt.Printf("Deleted todo with ID: %d\n", todoId)
}
