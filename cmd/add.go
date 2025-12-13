package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new todo",
	Long:  "Add a new todo with description and optional priority (high, medium, low)",
	Args:  cobra.MinimumNArgs(1),
	Run:   addTodo,
}

func addTodo(cmd *cobra.Command, args []string) {
	description := strings.Join(args, " ")

	err := todoList.Add(description)
	if err != nil {
		fmt.Printf("Error adding todo: %v\n", err)
		return
	}

	saveTodos()
	fmt.Printf("Added: %s\n", description)
}
