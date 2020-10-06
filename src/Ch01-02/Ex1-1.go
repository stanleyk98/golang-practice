// 1장 연습문제
// 커맨드라인 인수로 입력된 숫자들을 모두 더하는 go 프로그램 작성하기

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result int
	var count int
	var n int

	if len(os.Args) == 0 {
		fmt.Println("1개 이상의 입력값이 필요함")
		os.Exit(1)
	}
	arguments := os.Args

	for i := 1; i < len(arguments); i++ {
		n := strconv.ParseFloat(arguments[i], 64)
		result += n
		count++
	}
	fmt.Println("Total : ", result)

}
