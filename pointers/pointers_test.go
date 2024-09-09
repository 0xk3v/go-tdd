package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Testing Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Testing Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(200)}
		wallet.Withdraw(Bitcoin(50))
		assertBalance(t, wallet, Bitcoin(150))
	})

	t.Run("Testing Withdraw with insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(50)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assetError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func assetError(t testing.TB, got error, want string) {
	t.Helper()

	if got == nil {
		fmt.Print("Was expecting an error")
	}

	if got.Error() != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
