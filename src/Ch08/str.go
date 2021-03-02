package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	r := strings.NewReader("test")
	fmt.Println("r 의 길이는 ", r.Len())

	b := make([]byte, 1)

	for {
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
	s := strings.NewReader("This is an error~!\n")
	fmt.Println("r의 길이는 : ", s.Len())
	n, err := s.WriteTo(os.Stderr)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("읽은 내용에 따라 %d bytes 쓰기 완료\n", n)
}
