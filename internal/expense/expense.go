package expense

import (
	"time"
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
	expenses, err := ReadExpensesFromFile()
	if err != nil {
		return err
	}

	id := int64(len(expenses) + 1)

	expense := NewExpense(id, description, amount, category)

	expenses = append(expenses, *expense)

	err = WriteExpensesToFile(expenses)
	if err != nil {
		return err
	}

	return nil
}
