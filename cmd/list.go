package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

func ListCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "list",
		Short: "List Expenses",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunListExpensesCmd(args)
		},
	}

	// addCmd.Flags().StringVarP(&Description, "description", "d", "", "description of the expense")
	// addCmd.MarkFlagRequired("description")
	// addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "Amount of the expense")
	// addCmd.MarkFlagRequired("amount")
	// addCmd.Flags().StringVarP(&Category, "category", "c", "", "Category of the expense")
	//
	return addCmd
}

// func RunAddExpenseCmd(args []string) error {
// 	if Amount < 0 {
// 		return fmt.Errorf("Amount cannot be negative")
// 	}
//
// 	return expense.AddExpense(Description, Amount, Category)
// }

func RunListExpensesCmd(args []string) error {
	return expense.ListExpenses()
}
