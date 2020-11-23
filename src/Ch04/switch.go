package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("한문자 이상 입력 바람")
		os.Exit(1)
	}

	number, err := strconv.Atoi(arguments[1])

	if err != nil {
		fmt.Println("입력된 값은 int형이 아닙니다", number)
	} else {
		switch {
		case number < 0:
			fmt.Println("0 보다 작은 값~!")
		case number > 0:
			fmt.Println("0 보다 큰 값~!")
		default:
			fmt.Println("0 입니다~!")
		}
	}

	asString := arguments[1]
	switch asString {
	case "5":
		fmt.Println("다섯")
	case "0":
		fmt.Println("제로")
	default:
		fmt.Println("모름")
	}

	var negative = regexp.MustCompile(`-`)
	var floatingPoint = regexp.MustCompile(`\d?\. \d`)
	var email = regexp.MustCompile(`^[^@]+@[^@.]+\.[^@.]+`)

	switch {
	case negative.MatchString(asString):
		fmt.Println("Negative Number")
	case floatingPoint.MatchString(asString):
		fmt.Println("Floating Point")
	case email.MatchString(asString):
		fmt.Println("Email 주소입니다.")
		fallthrough
	default:
		fmt.Println("Something Else~!")
	}

	var aType error = nil
	switch aType.(type) {
	case nil:
		fmt.Println("Nil interface~!")
	default:
		fmt.Println("Nil interface 아님")
	}
}
