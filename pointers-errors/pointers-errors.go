package pointerserrors

import (
	"errors"
	"fmt"
)



type Bitcoin int

var ErrInsufficientBalance = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string{
  return fmt.Sprintf("%d BTC", b)
}


type Wallet struct {
 balance Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin){
  w.balance += deposit
}


func (w *Wallet) Balance() Bitcoin{
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error{
	if amount > w.balance {
		return ErrInsufficientBalance
	}
	w.balance -= amount
	return nil
}