package wallet

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // if lowercase then it's private and available only in package where it was defined
}

var ErrInsufficientFunds = errors.New("No enough money!")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
