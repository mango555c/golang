package pointers

import (
	"errors"
	"fmt"
)

// creat a new type: Bitcoin
type Bitcoin int

// create a strct: Wallet
type Wallet struct {
	balance Bitcoin
}

// create a new interface: Stringer
type Stringer interface {
	String() string
}

// define how the Bitcoin print
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// * means a pointer to a Wallet, o/w it's a copy
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	return w.balance
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
