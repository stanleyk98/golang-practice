package main

import (
	"flag"
	"fmt"
)

func main() {

	// 첫번째 인수로 bool 값을 받고 k 라고 명칭한다, 기본값은 true
	// 두번째 인수로 int 값을 받고 0 이라고 명칭한다, 기본값은 1
	minusK := flag.Bool("k", true, "k")
	minus0 := flag.Int("0", 1, "0")

	//커맨드라인 옵션에 대해 정의했으면 꼭 parse()함수를 호출할 것
	flag.Parse()

	// 플래그에 대해 입력된 값을 변수에 담을때 자동형변환 지원
	valueK := *minusK
	value0 := *minus0

	// value0 값 1 증가
	value0++

	fmt.Println("-k: ", valueK)
	fmt.Println("-0: ", value0)
}

// bool 값을 인수로 받을때는 '=' 으로 ?
// int 값을 인수로 받을때는 상관이 없나?

// go run simpleflag.go -k false -0 201		--> 왜 bool값은 false 가 되지않고 int값이 2가 출력되는가?
// go run simpleflag.go -k=false -0 201		--> 왜 int값이 202가 출력되는가?
// go run simpleflag.go -0= 100  			--> 띄어쓰면 에러출력

// go run simpleflag.go -k=false -0=201
// go run simpleflag.go -0 100
