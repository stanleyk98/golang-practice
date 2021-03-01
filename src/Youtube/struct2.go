package main

import (
	"fmt"
)

// 구조체 선언
type Student struct {
	name    string
	class   int
	sungjuk Sungjuk
}

// 구조체 타입 선언 (커스텀 타입도 구조체로 선언할수 있다)
type Sungjuk struct {
	name  string
	grade string
}

// 구조체 기능 선언 (과목 성적보여주기)
func (s *Student) ViewSungjuk() {
	fmt.Println(s.sungjuk)
}

// 입력받는 값에 따라 변동되는 값이라면 구조체 타입으로 받으면 안되고 포인터 타입으로 받아야 함
// 포인터를 호출하면 메모리 주소만 복사됨 (쓸데없는 메모리 낭비를 막을수 있다)
// 값 호출하면 전체 속성이 복사됨
func (s *Student) InputSungjuk(name string, grade string) {
	s.sungjuk.name = name
	s.sungjuk.grade = grade
}

func main() {

	var s Student

	s.name = "stanley"
	s.class = 1
	s.sungjuk.name = "English"
	s.sungjuk.grade = "A+"

	s.ViewSungjuk()

	// 포인터 타입으로 선언해놨기 때문에 값이 변경됨
	s.InputSungjuk("math", "B+")
	s.ViewSungjuk()

}
