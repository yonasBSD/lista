package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

var completeCmd = &cobra.Command{
	Use:   "complete [ID]",
	Short: "Complete a Todo",
	Long:  "Complete a Todo in the the todo list and mark status as completed",
	Args:  cobra.MinimumNArgs(1),
	Run:   completeTodo,
}

func completeTodo(cmd *cobra.Command, args []string) {
	todoId, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting argument into int")
		return
	}
	err = todoList.Complete(todoId)
	if err != nil {
		fmt.Printf("Error completing todo with id: %s", err)
		return
	}
}
