package wallet

import (
	"errors"
	"fmt"
)

type Egold int

func (b Egold) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	ballance Egold
}

func (w *Wallet) Balance() Egold {
	return w.ballance
}

func (w *Wallet) Deposit(amount Egold) {
	w.ballance += amount
}

var ErrInsufficientFunds = errors.New("insufficient funds for withdraw")

func (w *Wallet) Withdraw(amount Egold) error {
	if w.ballance < amount {
		return ErrInsufficientFunds
	}

	w.ballance -= amount
	return nil
}