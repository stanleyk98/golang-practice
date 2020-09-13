package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	argument := os.Args

	if len(arguments) == 1 {
		myString = "please give me one argument~!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
