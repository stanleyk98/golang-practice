package main

import (
	"fmt"
)

type Info struct {
	Name   string
	Mail   string
	Height int
	age    int
}

// pointer 를 사용하는 방식의 struct 생성하기
func createStruct(n, m string, h int, a int) *Info {
	if h > 200 {
		h = 0
	}
	return &Info{n, m, h, a}
}

// pointer 를 사용하지 않는 방식의 struct 생성하기
func retStruct(n, m string, h int, a int) Info {
	if h > 200 {
		h = 0
	}
	return Info{n, m, h, a}
}

func main() {
	s1 := createStruct("Stanley", "stanleyk@sk.com", 180, 300)
	s2 := retStruct("Terry", "test@sk.com", 190, 400)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)
}
