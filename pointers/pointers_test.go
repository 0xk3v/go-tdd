package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Testing Deposit", func(t *testing.T) {
		// Initializing wallet
		wallet := Wallet{}

		// Operations
		wallet.Deposit(10)
		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})

	t.Run("Testing Withdraw", func(t *testing.T) {
		// Setting up Wallets
		wallet := Wallet{}
		wallet.Deposit(200)

		err := wallet.Withdraw(50)
		if err != nil {
			t.Fatal(err)
		}

		got := wallet.Balance()
		want := 150

		if got != want {
			t.Errorf("Got %d, want %d", got, want)
		}
	})

	t.Run("Testing Withdraw without enough money", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(100)

		err := wallet.Withdraw(100)
		if err != nil {
			t.Fatal(err)
		}
	})
}
