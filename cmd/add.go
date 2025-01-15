package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/todevmilen/expense-tracker/internal/expense"

	_ "github.com/charmbracelet/lipgloss"
)

var Description string
var Amount float64
var Category string

func AddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Adds new expense",

		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddExpenseCmd(args)
		},
	}

	addCmd.Flags().StringVarP(&Description, "description", "d", "", "description of the expense")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "Amount of the expense")
	addCmd.MarkFlagRequired("amount")
	addCmd.Flags().StringVarP(&Category, "category", "c", "", "Category of the expense")

	return addCmd
}

func RunAddExpenseCmd(args []string) error {
	if Amount < 0 {
		return fmt.Errorf("Amount cannot be negative")
	}

	return expense.AddExpense(Description, Amount, Category)
}
