package dictionary

import "testing"

func TestDictionary(t *testing.T) {
	d := Dictionary{"test": "test def"}
	got := d.Search("test")
	want := "test def"
	assert(t, got, want)
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
