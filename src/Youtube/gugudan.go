package main

import (
	"fmt"
)

func main() {

	// 구구단 예제
	for dan := 1; dan <= 9; dan++ {

		// 5단은 제외하고 싶을때
		if dan == 5 {
			continue
		}

		fmt.Printf("%d단\n", dan)

		for j := 1; j <= 9; j++ {
			fmt.Printf("%d * %d = %d\n", dan, j, dan*j)
		}
		fmt.Println("-------------")
	}

	// 별 그리기 예제 (별 증가)
	for i := 0; i < 5; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	// 별 그리기 예제 (별 감소)
	for i := 0; i < 5; i++ {
		for j := 0; j < 5-i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
