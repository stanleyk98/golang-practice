package main

import (
	"fmt"
)

func Increase(x *int) {
	// x 의 메모리 주소에 있는 value 에 +1 을 하라
	*x = *x + 1
}

func main() {

	/*	포인터 예제 1
		var a int
		var b int
		var p *int // 포인터도 타입이다 포인터의 타입 표시는 *int / *string ....

		p = &a // 포인터 p 는 변수 a 의 메모리 주소다
		a = 3
		b = 2

		fmt.Println(a)
		fmt.Println(p)
		fmt.Println(*p) //포인터 p 의 메모리 주소에 있는 value 값이 *p 이다

		p = &b // 포인터 p 는 변수 b 의 메모리 주소다

		fmt.Println(b)
		fmt.Println(p)
		fmt.Println(*p) //포인터 p 의 메모리 주소에 있는 value 값이 *p 이다
	*/

	var a int
	a = 1

	Increase(&a)

	fmt.Println(a)

}
