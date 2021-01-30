package myInterface

type Shape interface {
	Area() float64
	Perimeter() float64
	// Area 함수와 Perimeter 함수를 구현하면 되며 두 함수 모두 float64 값 반환함
}
