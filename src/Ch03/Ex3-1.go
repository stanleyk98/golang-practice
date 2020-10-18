//연습문제 1
//숫자 4의 제곱에 대한 상수생성자 iota 를 작성한다
//일주일 요일에 대한 상수생성자 iota 를 작성한다

package main

import (
	"fmt"
)

// 4의 제곱 iota  (쉬프트연산자 사용)
const (
	a = 4 << iota
	_
	b
	_
	c
	_
	d
)

// 일주일의 요일 iota
const (
	Sunday = iota + 1
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	// 4의 제곱 iota 출력
	fmt.Println("4의 1제곱: ", a)
	fmt.Println("4의 2제곱: ", b)
	fmt.Println("4의 3제곱: ", c)
	fmt.Println("4의 4제곱: ", d)

	// 일주일의 요일 상수 값 출력
	fmt.Println("일요일: ", Sunday)
	fmt.Println("수요일: ", Wednesday)
	fmt.Println("금요일: ", Friday)
}
