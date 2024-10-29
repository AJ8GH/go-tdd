package countdown

import (
	"bytes"
	"fmt"
	"testing"
)

type SpySleeper struct {
	calls int
}

func (s *SpySleeper) Sleep() {
	s.calls++
}

func TestCountdown(t *testing.T) {
	buffer := bytes.Buffer{}
	sleeper := SpySleeper{}
	Countdown(&sleeper, &buffer)
	got := buffer.String()
	want := "3\n2\n1\nGo!"
	assert(t, got, want, "%q")

	sleepsGot := sleeper.calls
	sleepsWant := 3
	assert(t, sleepsGot, sleepsWant, "%d")
}

func assert[T comparable](t testing.TB, got, want T, verb string) {
	if got != want {
		t.Errorf(fmt.Sprintf("got %v want %v", verb, verb), got, want)
	}
}
