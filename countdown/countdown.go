package countdown

import (
	"fmt"
	"io"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
)

func Countdown(w io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWord)
}
