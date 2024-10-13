package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("No args", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("With name", func(t *testing.T) {
		got := Hello("Adam", "English")
		want := "Hello, Adam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In Spanish", func(t *testing.T) {
		got := Hello("Adam", "Spanish")
		want := "Hola, Adam"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In French", func(t *testing.T) {
		got := Hello("Adam", "French")
		want := "Bonjour, Adam"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
