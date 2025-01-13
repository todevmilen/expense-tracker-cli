package expense

import (
	"os"

	"github.com/todevmilen/expense-tracker/internal/log"
)

func expensesFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Error("Error getting current working directory: ", err)
	}
}

func ReadExpensesFromFile() ([]Expense, error) {

}
