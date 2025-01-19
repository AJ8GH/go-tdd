package greet

import (
	"fmt"
	"io"
)

const greeting = "Hello, %s"

func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, greeting, name)
}
