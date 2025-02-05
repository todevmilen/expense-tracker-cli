/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import "github.com/spf13/cobra"

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "Expense Tracker is a CLI tool to manage your expenses",
		Long: `Manage your expenses with ease using this CLI tool ! 
		#------------------------Expense Tracker---------------------------#
	`,
	}

	cmd.AddCommand(HelloCmd())
	cmd.AddCommand(AddCmd())
	cmd.AddCommand(ListCmd())
	cmd.AddCommand(DeleteCmd())
	cmd.AddCommand(SummaryCmd())
	cmd.AddCommand(UpdateCmd())

	return cmd
}
