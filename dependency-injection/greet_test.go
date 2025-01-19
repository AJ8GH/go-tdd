package greet

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Adam")

	got := buffer.String()
	want := "Hello, Adam"

	assert(t, got, want)
}

func assert(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
