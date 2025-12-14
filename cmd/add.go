package cmd

import (
	"fmt"
	"strings"

	"github.com/kwame-Owusu/lista/internal/models"
	"github.com/spf13/cobra"
)

var priorityFlag string

var addCmd = &cobra.Command{
	Use:   "add [description]",
	Short: "Add a new todo",
	Long:  "Add a new todo with description and optional priority (high, medium, low)",
	Args:  cobra.MinimumNArgs(1),
	Run:   addTodo,
}

func init() {
	addCmd.Flags().StringVarP(&priorityFlag, "priority", "p", "low", "Priority level (high/h, medium/m, low/l)")
}

func addTodo(cmd *cobra.Command, args []string) {
	description := strings.Join(args, " ")
	// Parse the priority flag
	priority, err := models.ParsePriority(priorityFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = todoList.Add(description, priority)
	if err != nil {
		fmt.Printf("Error adding todo: %v\n", err)
		return
	}

	saveTodos()
	fmt.Printf("Added: %s\n", description)
}
