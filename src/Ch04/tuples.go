package main

import (
	"fmt"
)

func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
	//2를 곱한값 , 제곱값, -값
}

func main() {
	fmt.Println(retThree(10))

	n1, n2, n3 := retThree(20)
	fmt.Println(n1, n2, n3)

	n1, n2 = n2, n1 // 값 스왑
	fmt.Println(n1, n2, n3)

	x1, x2, x3 := n1*2, n1*n1, -n1
	fmt.Println(x1, x2, x3)

}
