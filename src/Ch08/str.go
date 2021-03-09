package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	r := strings.NewReader("stanley")
	fmt.Println("r 의 길이는 ", r.Len())

	// byte 슬라이스 생성
	b := make([]byte, 1)

	for {
		// byte 단위로 string 읽기함수 -> read()
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("읽은 문자열 %s 는 %d Bytes 입니다. \n", b, n)
	}
	s := strings.NewReader("Go error~!\n")
	fmt.Println("r의 길이는 : ", s.Len())

	// 표준에러에 출력하기 -> writeto()
	n, err := s.WriteTo(os.Stderr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("표준에러로 %d bytes 쓰기 완료\n", n)
}

// go run str.go
