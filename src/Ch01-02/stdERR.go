package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args

	if len(arguments) == 1 {
		myString = "하나이상의 인수를 입력하세요"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, "이건 표준출력\n")
	io.WriteString(os.Stderr, myString)
	io.WriteString(os.Stderr, "\n")
}
