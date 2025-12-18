package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

var addNotesCmd = &cobra.Command{
	Use:   "notes [id]",
	Short: "Adds notes to a todo",
	Long:  "Adds notes to a todo, if not present. if present it appends to the existing note",
	Args:  cobra.MinimumNArgs(1),
	Run:   addNotes,
}

func addNotes(cmd *cobra.Command, args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting id to int: %v", err)
		return
	}

	todo, err := todoList.GetByID(id)
	if err != nil {
		fmt.Printf("No todo with ID %d exists", id)
		return
	}
	newNotes := strings.Join(args[1:], " ")

	if todo.ID == id {
		if todo.Notes != "" {
			todo.Notes += " " + newNotes
		} else {
			todo.Notes = newNotes
		}
		saveTodos()
		return
	}
	fmt.Printf("No todo with ID %v exists\n", id)
}
