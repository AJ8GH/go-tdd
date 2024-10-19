package pointers

import "testing"

func TestWallet(t *testing.T) {
	w := Wallet{}
	w.Deposit(10)

	got := w.Balance()
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
