package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("하나이상의 인수를 입력하세요")
		// 프로그램 실행종료
		os.Exit(1)
	}

	// 이번 예제에만 적용할 strconv.parsefloat 에러리턴값을 무시하기 위한 임의 코딩
	arguments := os.Args

	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)

	for i := 2; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)

		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Min : ", min)
	fmt.Println("Max : ", max)
}
