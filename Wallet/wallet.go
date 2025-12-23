package Wallet

import "errors"

type Payment interface {
	Pay(amount float64) error
}

type Wallet struct {
	Balance float64
}

// Add money
func (w *Wallet) Deposit(amount float64) {
	w.Balance += amount
}

// Pay money
func (w *Wallet) Pay(amount float64) error {
	if amount > w.Balance {
		return errors.New("insufficient funds")
	}
	w.Balance -= amount
	return nil
}
