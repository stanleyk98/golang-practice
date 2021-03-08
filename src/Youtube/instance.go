package main

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	grade int
}

// 객체 지형의 매소드 정의
// Student 구조체를 포인터로 받는 t 가 주체가 되면서 앞으로 나가면 a. 의 객체로 사용가능하다
func (t *Student) SetName(newName string) {
	t.name = newName
}

// 프로시저 지향의 함수 정의
func SetName(t *Student, name string) {
	t.name = name
}

func main() {
	var a *Student

	a = &Student{"aaa", 20, 10}

	// SetName 함수를 호출하든, a객체의 SetName 메소드를 호출하든 결과는 같지만
	// 객체지향적인 매소드 호출하는 것이 더욱 편하다
	// a 객체를 만든다 = 인스턴스

	SetName(a, "bbb")
	a.SetName("bbb")

	// a.name = "bbb"

	fmt.Println(a)
}
