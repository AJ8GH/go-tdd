package countdown

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
	"time"
)

const (
	write = "write"
	sleep = "sleep"
)

type SpyCountdownOperations struct {
	Calls []string
}

func (c *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

func (c *SpyCountdownOperations) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("happy", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Countdown(&SpyCountdownOperations{}, &buffer)
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		assert(t, got, want, "%q")
	})

	t.Run("order", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		got := spySleepPrinter.Calls

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	assert(t, spyTime.durationSlept, sleepTime, "%q")
}

func assert[T comparable](t testing.TB, got, want T, verb string) {
	t.Helper()
	if got != want {
		t.Errorf(fmt.Sprintf("got %v want %v", verb, verb), got, want)
	}
}
