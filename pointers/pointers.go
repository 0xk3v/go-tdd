package pointers

type Wallet struct {
	amount int
}

func (w Wallet) Deposit(amount int) {
	w.amount = w.amount + amount
}

func (w Wallet) Balance() int {
	return w.amount
}
