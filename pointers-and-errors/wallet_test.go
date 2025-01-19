package wallet

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
		assertError(t, err, nil)
		assert(t, w, 5)
	})

	t.Run("withdraw error", func(t *testing.T) {
		w := Wallet{10}
		err := w.Withdraw(500)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assert(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
