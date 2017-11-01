package main

import (
	"theduke/repl"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
