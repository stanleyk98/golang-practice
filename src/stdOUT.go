package main

import (
	//표준출력 os.Stdout 은 io 패키지를 사용한다
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "please give me one argument~!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}
