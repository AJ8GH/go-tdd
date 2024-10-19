package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {

		w := Wallet{}
		w.Deposit(10)

		got := w.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {

		w := Wallet{10}
		w.Withdraw(5)

		got := w.Balance()
		want := Bitcoin(5)

		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	})
}
