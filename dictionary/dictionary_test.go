package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	t.Run("happy get", func(t *testing.T) {
		d := Dictionary{"test": "test def"}
		got, err := d.Search("test")
		assert(t, err, nil)
		assert(t, got, "test def")
	})

	t.Run("sad get", func(t *testing.T) {
		d := Dictionary{}
		got, err := d.Search("test")
		assert(t, err, ErrNoWordFound)
		assert(t, got, "")
	})

	t.Run("happy add", func(t *testing.T) {
		d := Dictionary{}
		addErr := d.Add("test", "test def")
		assert(t, addErr, nil)

		got, searchErr := d.Search("test")
		assert(t, searchErr, nil)
		assert(t, got, "test def")
	})

	t.Run("sad add", func(t *testing.T) {
		d := Dictionary{}
		d.Add("test", "test def")
		addErr := d.Add("test", "new test def")
		assert(t, addErr, ErrWordExists)

		got, searchErr := d.Search("test")
		assert(t, got, "test def")
		assert(t, searchErr, nil)
	})

	t.Run("happy update", func(t *testing.T) {
		d := Dictionary{"test": "test def"}
		updateErr := d.Update("test", "new test def")
		assert(t, updateErr, nil)

		got, searchErr := d.Search("test")
		assert(t, got, "new test def")
		assert(t, searchErr, nil)
	})

	t.Run("sad update", func(t *testing.T) {
		d := Dictionary{}
		updateErr := d.Update("test", "new test def")
		assert(t, updateErr, ErrNoWordFound)
	})

	t.Run("happy delete", func(t *testing.T) {
		d := Dictionary{"test": "test def"}
		deleteErr := d.Delete("test")
		assert(t, deleteErr, nil)

		got, searchErr := d.Search("test")
		assert(t, got, "")
		assert(t, searchErr, ErrNoWordFound)
	})

	t.Run("sad delete", func(t *testing.T) {
		d := Dictionary{}
		deleteErr := d.Delete("test")
		assert(t, deleteErr, ErrNoWordFound)
	})
}

func assert[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
