package main

import (
	"fmt"
)

func main() {

	// struct 는 보통 전역범위에서 사용하기 때문에 main함수 밖에서 선언하지만 함수 안에서 로컬범위로 사용해도 상관은 없음

	type XYZ struct {
		X int
		Y int
		Z int
	}
	var s1 XYZ // XYZ 구조체를 담는 s1 변수 생성

	fmt.Println(s1.Y, s1.Z) // struct 내의 값 접근

	p1 := XYZ{300, 12, -2}
	p2 := XYZ{Z: 12, Y: 13, X: 300} // struct 내의 순서에 상관없이 값 대입가능
	fmt.Println(p1)
	fmt.Println(p2)

	pSlice := [4]XYZ{} //4원소 배열 생성
	pSlice[2] = p1     //인덱스2에 p1
	pSlice[0] = p2     //인덱스0에 p2
	//인덱스 1,4는 할당값 없으므로 기본값 0 유지
	fmt.Println(pSlice)

	p2 = XYZ{1, 2, 3} // struct의 배열값을 바꿔도 원본 struct 의 값에는 영향없음
	fmt.Println(pSlice)

}
