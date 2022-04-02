package main

import (
	"os"

	"github.com/regmicmahesh/small-lang/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
