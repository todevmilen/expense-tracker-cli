package cmd

import (
	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

func UpdateCmd() *cobra.Command {
	updateCmd := &cobra.Command{
		Use:   "update",
		Short: "update expense by id",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunUpdateExpenseCmd(args)
		},
	}

	updateCmd.Flags().Int64VarP(&Id, "id", "i", 0, "id of the expense")
	updateCmd.MarkFlagRequired("id")
	updateCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "amount of the expense")
	updateCmd.MarkFlagRequired("amount")

	return updateCmd
}

// func RunAddExpenseCmd(args []string) error {
// 	if Amount < 0 {
// 		return fmt.Errorf("Amount cannot be negative")
// 	}
//
// 	return expense.AddExpense(Description, Amount, Category)
// }

func RunUpdateExpenseCmd(args []string) error {
	return expense.UpdateExpense(Id, Amount)
}
