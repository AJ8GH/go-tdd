package sync

import (
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing single thread", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		want := 3
		got := counter.Value()

		if want != got {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
