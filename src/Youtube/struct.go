package main

import (
	"fmt"
)

// 구조체의 선언
type Person struct {
	name string
	age  int
}

// 구조체의 기능 추가 - 이름출력하는 기능
func (p Person) printName() {
	fmt.Println(p.name)
}

func main() {
	var p Person

	p1 := Person{"땡그리", 24}
	p2 := Person{name: "살구", age: 25}
	p3 := Person{name: "stanley"}
	p4 := Person{}

	fmt.Println(p, p1, p2, p3, p4)

	// 구조체의 속성값에 입력
	p.name = "won"
	p.age = 40

	// 구조체의 메소드 호출
	p.printName()

}
