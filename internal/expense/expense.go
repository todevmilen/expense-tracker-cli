package expense

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/todevmilen/expense-tracker/internal/log"
)

type Expense struct {
	ID          int64     `json:"id"`
	Description string    `json:"description"`
	Amount      float64   `json:"amount"`
	Category    string    `json:"category"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewExpense(id int64, description string, amount float64, category string) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		Category:    category,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func AddExpense(description string, amount float64, category string) error {
	var id int64
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		id = 1
	} else {
		id = int64(expenses[len(expenses)-1].ID) + 1
	}

	expense := NewExpense(id, description, amount, category)

	expenses = append(expenses, *expense)

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	log.Success(fmt.Sprintf("Expense added successfully (ID: %v)", id))

	return nil
}

func ListExpenses(category string) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	fmt.Printf("%-5s %-12s %-20s %-10s %-10s\n", "ID", "Date", "Description", "Amount", "Category")
	for _, expense := range expenses {
		if category != "" {
			if expense.Category != category {
				continue
			}
		}

		fmt.Printf("%-5d %-12s %-20s %-10.2f %-10s\n",
			expense.ID,
			expense.CreatedAt.Format("02.01.2006"),
			expense.Description,
			expense.Amount,
			expense.Category,
		)
	}

	return nil
}

func DeleteExpense(id int64) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	if len(expenses) == 0 {
		return errors.New("Cannot delete! There are no expenses")
	}

	updatedExpenses := slices.DeleteFunc(expenses, func(e Expense) bool {
		return e.ID == id
	})

	err = WriteExpensesToFile(updatedExpenses)
	if err != nil {
		return err
	}

	log.Success(fmt.Sprintf("Expense with ID: %v was deleted!", id))

	return nil
}

func Summary() error {
	var summaryValue float64

	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	for _, expense := range expenses {
		summaryValue += expense.Amount
	}

	log.Info(fmt.Sprintf("Total expenses: $%.2f", summaryValue))

	return nil
}

func UpdateExpense(id int64, amount float64) error {
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}
	if len(expenses) == 0 {
		return errors.New("Cannot update! There are no expenses")
	}

	expenseIndex := slices.IndexFunc(expenses, func(e Expense) bool {
		return e.ID == id
	})

	expenses[expenseIndex].Amount = amount

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	log.Success(fmt.Sprintf("Expense with ID: %v was updated with amount of: %.2f", id, amount))

	return nil
}
