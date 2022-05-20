package entities

import (
	"fmt"
	"time"
)

type Account struct {
	ID        string
	UserID    string
	Balance   int
	Limit     int
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AccountProps struct {
	ID     string
	UserID string
}

func NewAccount(props AccountProps) *Account {
	return &Account{
		ID:        props.ID,
		UserID:    props.UserID,
		Balance:   0,
		Status:    "PENDING",
		Limit:     500000,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (account *Account) Deposit(amount int) error {
	if amount > account.Limit && time.Now().Weekday() == 0 {
		return fmt.Errorf("Sorry you have exceeded your limit for Sunday")
	}

	account.Balance += amount
	return nil
}

func (account *Account) Withdraw(amount int) error {
	if amount > account.Balance {
		return fmt.Errorf("Insufficient funds")
	}

	account.Balance -= amount
	return nil
}
