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
		err := wallet.Withdraw(Bitcoin(50))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(150))
	})

	t.Run("Testing Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(100))

		assetError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(50))
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Got an error but didn't expect one")
	}
}

func assetError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		fmt.Print("Was expecting an error")
	}

	if got != want {
		t.Errorf("Got %q, wanted %q", got, want)
	}
}
