package main

import (
	"fmt"
)

type square struct {
	X float64
}

type circle struct {
	R float64
}

type rectangle struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case square:
		fmt.Println("구형")
	case circle:
		fmt.Printf("%v 는 원형\n", v)
	case rectangle:
		fmt.Println("이건 직사각형")
	default:
		fmt.Printf("알려지지않은 타입 %T!\n", v)
	}
}

func main() {
	x := circle{R: 10}
	tellInterface(x)

	y := rectangle{X: 4, Y: 1}
	tellInterface(y)

	z := square{X: 4}
	tellInterface(z)

	tellInterface("stanlely")
}
