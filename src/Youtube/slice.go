package main

import (
	"fmt"
)

func main() {

	/*
		a := make([]int, 0, 8)

		fmt.Printf("len(a) : %d\n", len(a))
		fmt.Printf("cap(a) : %d\n", cap(a))

		a = append(a, 100)
		fmt.Printf("len(a) : %d\n", len(a))
		fmt.Printf("cap(a) : %d\n", cap(a))
	*/

	/* 슬라이스를 다른 슬라이스로 대입하려면 메모리 주소를 잘 봐야함 (본래의 슬라이스인지 복사된 신규 슬라이스인지)
	a := make([]int, 2, 2)

	a[0] = 1
	a[1] = 2

	// 슬라이스 a 가 여유공간이 없으므로 cap 2배인 슬라이스 b 를 생성하고 거기에 b 원소가 들어간다
	// 슬라이스 a 와 슬라이스 b 는 다른 메모리 주소를 가진다
	b := append(a, 100)

	b[0] = 101
	b[1] = 201

	fmt.Println(a)
	fmt.Println(b)
	*/

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// a 의 인덱스 4부터 8까지의 슬라이스를 b 슬라이스로 참고한다
	// a 와 b 는 같은 메모리 주소
	// start index ~ enc index까지 단, end index 는 포함하지 않는다
	b := a[4:8]

	// a 의 인덱스 4부터 끝까지 슬라이스를 c 슬라이스로 참고한다
	// a 와 c 는 같은 메모리 주소
	c := a[4:]

	// a 의 인덱스 처음부터 end index 앞까지 슬라이스를 d 슬라이스로 참고한다
	// a 와 d 는 같은 메모리 주소
	d := a[:8]

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}
