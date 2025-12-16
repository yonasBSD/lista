package cmd

// import (
// 	"fmt"
// 	"github.com/charmbracelet/bubbletea"
// 	"github.com/spf13/cobra"
// 	"os"
// )
//
// var tuiCmd = &cobra.Command{
// 	Use:   "tui",
// 	Short: "Run lista in interactive TUI mode",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		p := tea.NewProgram(tui.NewModel(tui.ModeList))
// 		if _, err := p.Run(); err != nil {
// 			fmt.Printf("an error occurred: %v", err)
// 			os.Exit(1)
// 		}
// 	},
// }
