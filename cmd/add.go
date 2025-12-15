package cmd

import (
	"fmt"
	"strings"

	"github.com/kwame-Owusu/lista/internal/models"
	"github.com/spf13/cobra"
)

var priorityFlag string
var notesFlag string

var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo",
	Long:  "Add a new todo with description and optional priority (high, medium, low) and notes",
	Args:  cobra.MinimumNArgs(1),
	Run:   addTodo,
}

func init() {
	addCmd.Flags().StringVarP(&priorityFlag, "priority", "p", "low", "Priority level (high/h, medium/m, low/l)")
	addCmd.Flags().StringVarP(&notesFlag, "notes", "n", "", "notes (lorem ipsum)")
}

func addTodo(cmd *cobra.Command, args []string) {
	title := strings.Join(args, " ")
	// Parse the priority flag
	priority, err := models.ParsePriority(priorityFlag)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	notes := notesFlag

	err = todoList.Add(title, priority, notes)
	if err != nil {
		fmt.Printf("Error adding todo: %v\n", err)
		return
	}

	saveTodos()
}
