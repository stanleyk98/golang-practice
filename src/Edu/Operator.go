package main

import (
	"fmt"
)

func main() {
	a := 2
	var num int

	num = a
	fmt.Println("num = a : ", num)

	num += 4
	fmt.Println("num += 4 : ", num)

	//다른 예제
	var num1, num2, num3, result int

	fmt.Println("3자리 입력하기")
	fmt.Println("천천히 입력해주세요~")
	fmt.Scanln(&num1, &num2, &num3)

	//result = &num1 x &num2 + &num3

	fmt.Printf("%d x %d + %d = %d\n", num1, num2, num3, result)

}
