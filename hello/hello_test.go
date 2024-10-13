package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("no args", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("With name", func(t *testing.T) {
		got := Hello("Adam")
		want := "Hello, Adam"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
