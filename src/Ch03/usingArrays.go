package main

import (
	"fmt"
)

func main() {
	anArray := [4]int{1, 2, 4, -4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {7, 0}}}

	fmt.Println(anArray, "의 길이는 ", len(anArray))
	fmt.Println(twoD, "의 첫번째 원소는 ", twoD[0][0])
	fmt.Println(threeD, "의 길이는 ", len(threeD))

	for i := 0; i < len(threeD); i++ { // len(threeD) 가 2 이므로 i는 0,1
		v := threeD[i] // threeD[0] 은 (1,0)(-2,4) 이고 threeD[1] 은 (5,-1)(7,0) 이다

		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}
	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println()
	}
}
