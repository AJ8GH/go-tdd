package main

import (
	"os"

	"github.com/aj8gh/gotdd/greet"
)

func main() {
	greet.Greet(os.Stdout, "Adam")
}
