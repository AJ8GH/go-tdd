package countdown

import (
	"fmt"
	"io"
	"time"
)

const (
	countdownStart        = 3
	countdownPauseSeconds = 1
	finalWord             = "Go!"
)

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	Duration  time.Duration
	SleepFunc func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFunc(c.Duration)
}

func Countdown(s Sleeper, w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}
