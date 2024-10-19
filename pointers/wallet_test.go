package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(10)
		assert(t, w, 10)
	})

	t.Run("withdraw", func(t *testing.T) {
		w := Wallet{10}
		err := w.Withdraw(5)
		if err != nil {
			t.Errorf("error should be nil but got %q", err)
		}
		assert(t, w, 5)
	})

	t.Run("withdraw error", func(t *testing.T) {
		w := Wallet{10}
		err := w.Withdraw(500)
		if err == nil {
			t.Errorf("wanted error but got nil")
		}
		assert(t, w, 10)
	})
}

func assert(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}