package main

import (
	"fmt"
)

func main() {

	a := 4

	if a == 3 {
		fmt.Println("a 는 3")
	} else if a == 4 {
		fmt.Println("a 는 4")
	} else {
		fmt.Println("a 는 3도 아니고 4도 아님")
	}

	fmt.Println("a 의 값은 ", a)
}
