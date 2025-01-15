package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

func ListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List Expenses",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListExpensesCmd(args)
		},
	}

	listCmd.Flags().StringVarP(&Category, "category", "c", "", "Category of the expense")

	return listCmd
}

// func RunAddExpenseCmd(args []string) error {
// 	if Amount < 0 {
// 		return fmt.Errorf("Amount cannot be negative")
// 	}
//
// 	return expense.AddExpense(Description, Amount, Category)
// }

func RunListExpensesCmd(args []string) error {
	return expense.ListExpenses(Category)
}
