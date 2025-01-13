package main

import (
	"fmt"

	"github.com/todevmilen/expense-tracker/cmd"
)

func main() {
	rootCmd := cmd.RootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
