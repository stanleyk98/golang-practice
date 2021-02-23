package main

import (
	"fmt"
)

// 정수 인자값 2개를 받고 1개의 정수형으로 출력
func add(x int, y int) int {
	return x + y
}

func main() {

	for i := 0; i < 10; i++ {
		fmt.Printf("%d + %d = %d\n", i, i+2, add(i, i+2))
		
}
