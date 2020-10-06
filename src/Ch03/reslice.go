package main

import (
	"fmt"
)

func main() {
	s1 := make([]int, 5) // make 함수로 슬라이스 생성, int형으로 5개 이상의 원소를 가진 슬라이스가 만들어짐
	reslice := s1[1:3]   // s1 슬라이스의 두번째, 세번째 원소를 선택함 (0부터 시작)

	fmt.Println(s1)
	fmt.Println(reslice)

	reslice[0] = -100   // 첫번째 원소 접근
	reslice[1] = 123456 // 두번째 원소 접근

	fmt.Println(s1)
	fmt.Println(reslice)

}
