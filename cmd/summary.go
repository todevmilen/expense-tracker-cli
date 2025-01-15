package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

func SummaryCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "summary",
		Short: "Get summary of expenses amount",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummaryExpenses(args)
		},
	}

	listCmd.Flags().StringVarP(&Category, "category", "c", "", "Category of the expense")

	return listCmd
}

func RunSummaryExpenses(args []string) error {
	return expense.Summary()
}
