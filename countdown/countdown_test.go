package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	Countdown(&buffer)
	got := buffer.String()
	want := "3\n2\n1\nGo!"
	assert(t, got, want)
}

func assert(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
