package main

import (
	"fmt"
)

func main() {

	var s []int

	s = make([]int, 3)

	s[0] = 100
	s[1] = 200
	s[2] = 300

	// 기존 cap 3 인데 4개가 더 추가되어야 하므로 cap 을 늘린다
	//s = append(s, 40, 50, 60, 70)
	fmt.Println("------------------------")

	var t []int

	t = append(s, 400)

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))

	fmt.Println("------------------------")
	var u []int

	u = append(t, 500)

	u[0] = 9999
	// t 슬라이스에는 여유공간이 있으므로 같은 메모리 공간을 공유하므로 u 의 원소값을 바꾸면 t 의 원소값도 바뀐다

	fmt.Println(s, len(s), cap(s))
	fmt.Println(t, len(t), cap(t))
	fmt.Println(u, len(u), cap(u))
}
