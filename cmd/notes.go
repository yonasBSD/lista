package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var addNotesCmd = &cobra.Command{
	Use:   "notes [id]",
	Short: "adds notes to a todo",
	Long:  "adds notes to a todo, if not present. if present it appends to the existing note",
	Args:  cobra.MinimumNArgs(1),
	Run:   addNotes,
}

func addNotes(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting id to int: %v", err)
		return
	}

	todos := todoList.List()
	newNotes := strings.Join(args[1:], " ")

	for i := range todos {
		if todos[i].ID == id {
			if todos[i].Notes != "" {
				todos[i].Notes += " " + newNotes
			} else {
				todos[i].Notes = newNotes
			}
			saveTodos()
			return
		}
	}
	fmt.Printf("No todo with ID %v exists\n", id)
}
