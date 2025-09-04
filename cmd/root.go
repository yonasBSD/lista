package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command
var rootCmd = &cobra.Command{
	Use:   "todo-cli",
	Short: "A minimal todo CLI in Go",
	Long:  `todo-cli is a simple and aesthetic CLI app to manage your todos.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to todo-cli! Use 'todo-cli help' to see available commands.")
	},
}

// Execute executes the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		// Exit with error code 1 if something goes wrong
		// os.Exit(1)  // optional
	}
}
