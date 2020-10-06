package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("한자리 이상의 문자를 넣으시오")
		os.Exit(1)
	}

	arguments := os.Args
	var err error = errors.New("에러났다!!!!")
	k := 1
	var n float64

	for err != nil {
		if k >= len(arguments) {
			fmt.Println("float type 을 입력하시오")
			return
		}
		n, err = strconv.ParseFloat(arguments[k], 64)
		k++
	}
	min, max := n, n

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if n < min {
				min = n
			}

			if n > max {
				max = n
			}
		}
	}
	fmt.Println("Min : ", min)
	fmt.Println("Max : ", max)

}
