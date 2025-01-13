package cmd

import (
	"github.com/spf13/cobra"

	"github.com/charmbracelet/lipgloss"
)

func HelloCmd() *cobra.Command {
	helloCmd := &cobra.Command{
		Use:   "hello",
		Short: "Prints hello world",
		Run: func(cmd *cobra.Command, args []string) {
			style := lipgloss.NewStyle().Foreground(lipgloss.Color("#FF00FF")).Background(lipgloss.Color("#0000FF")).Bold(true)
			cmd.Println(style.Render("Hello, World!"))
		},
	}

	return helloCmd
}
