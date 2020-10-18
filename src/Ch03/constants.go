package main

import (
	"fmt"
)

type Digit int
type Power2 int

const PI = 3.1415926

const (
	C1 = "C1C1C1"
	C2 = "C2C2C2"
	C3 = "C3C3C3"
)

func main() {
	const s1 = 123
	var v1 float32 = s1 * 12

	fmt.Println(v1)
	fmt.Println(PI)

	// iota 로 1씩 값 증가
	const (
		Zero Digit = iota
		One
		Two
		Three
		Four
	)

	fmt.Println(One)
	fmt.Println(Two)

	const (
		p2_0 Power2 = 1 << iota
		_
		p2_2
		_
		p2_4
		_
		p2_6
	)

	fmt.Println("2의 0제곱: ", p2_0)
	fmt.Println("2의 2제곱: ", p2_2)
	fmt.Println("2의 4제곱: ", p2_4)
	fmt.Println("2의 6제곱: ", p2_6)
}
