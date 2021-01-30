package main

import (
	"fmt"
)

type twoInts struct {
	X int64
	Y int64
}

func regularFunction(a, b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func (a twoInts) method(b twoInts) twoInts {
	temp := twoInts{X: a.X + b.X, Y: a.Y + b.Y}
	return temp
}

func main() {
	i := twoInts{X: 1, Y: 2}
	j := twoInts{X: -5, Y: -2}

	fmt.Println(regularFunction(i, j))
	// regularfunction 에서 X값은 1 + -5 , Y값은 2 + -2   -> {-4 , 0}
	// -4, 0 으로 출력 안되고 왜 {-4,0} 으로 출력될까?

	fmt.Println(i.method(j))

}
