package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

var Id int64

func DeleteCmd() *cobra.Command {
	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete expense by id",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunDeleteExpenseCmd(args)
		},
	}

	deleteCmd.Flags().Int64VarP(&Id, "id", "i", 0, "id of the expense")

	return deleteCmd
}

// func RunAddExpenseCmd(args []string) error {
// 	if Amount < 0 {
// 		return fmt.Errorf("Amount cannot be negative")
// 	}
//
// 	return expense.AddExpense(Description, Amount, Category)
// }

func RunDeleteExpenseCmd(args []string) error {
	return expense.DeleteExpense(Id)
}
