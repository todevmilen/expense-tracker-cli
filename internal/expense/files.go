package expense

import (
	"encoding/json"
	"os"
	"path"

	"github.com/todevmilen/expense-tracker/internal/log"
)

func expensesFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.ErrorWithValue("Error getting current working directory: ", err)
	}

	return path.Join(cwd, "expenses.json")
}

func ReadExpensesFromFile() ([]Expense, error) {
	filePath := expensesFilePath()

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		log.Info("File does not exist. Creating file...")

		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			return nil, err
		}

		defer file.Close()

		return []Expense{}, nil
	}

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	expenses := []Expense{}
	err = json.NewDecoder(file).Decode(&expenses)

	if err != nil {
		return nil, err
	}

	return expenses, nil

}

func WriteExpensesToFile(expenses []Expense) error {
	filePath := expensesFilePath()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	err = json.NewEncoder(file).Encode(&expenses)
	if err != nil {
		return err
	}

	return nil
}
