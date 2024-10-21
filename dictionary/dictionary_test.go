package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		d := Dictionary{"test": "test def"}
		got, err := d.Search("test")
		assert(t, err, nil)
		assert(t, got, "test def")
	})

	t.Run("sad", func(t *testing.T) {
		d := Dictionary{}
		got, err := d.Search("test")
		assert(t, err, ErrNoWordFound)
		assert(t, got, "")
	})
}

func assert[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
