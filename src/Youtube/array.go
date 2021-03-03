package main

import "fmt"

func main() {

	// 배열 테스트 1번
	// var A [10]int
	// for i := 0; i < len(A); i++ {
	// 	A[i] = i * i
	// }
	// fmt.Println(A)

	//배열 테스트 2번
	// s := "Hello World"
	// s1 := "Hello 월드"

	// 한글을 문자열 배열로 표시할때 사용할수 있는 타입은 rune
	// s2 := []rune(s1)
	// fmt.Println("len(s1)= ", len(s1))
	// fmt.Println("len(s2)= ", len(s2))

	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(string(s[i]), ", ")
	// }
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Print(string(s2[i]), ", ")
	// }

	// s := []rune("stanley 김태욱")

	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(string(s[i]), ", ")
	// }

	// Redix 정렬
	arr := [11]int{0, 5, 4, 9, 1, 2, 8, 3, 6, 4, 5}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i]
		temp[idx]++
	}
	idx := 0
	for i := 0; i < len(temp); i++ {
		for j := 0; j < temp[i]; j++ {
			arr[idx] = i
			idx++
		}
	}
	fmt.Println(arr)
}
