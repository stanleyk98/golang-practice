package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 파일쓰는 5가지 방법

func main() {
	// byte 슬라이스
	s := []byte("파일 내용입니다~ \n")

	// 파일쓰기 방법1 : Fprintf() 함수 사용
	f1, err := os.Create("f1.txt")
	if err != nil {
		fmt.Println("파일을 생성할수 없음", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	// 파일쓰기 방법2 : WriteString() 함수 사용
	f2, err := os.Create("f2.txt")
	if err != nil {
		fmt.Println("파일을 생성할수 없음", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)

	// 파일쓰기 방법3 : bufio.NewWriter() 함수로 쓸파일 열고 bufio.WriteString() 함수로 데이터 쓰기
	f3, err := os.Create("f3.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	w := bufio.NewWriter(f3)
	n, err = w.WriteString(string(s))
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	// 파일쓰기 방법4 : ioutil.WriteFile() 함수 사용
	// ioutil 은 전체파일을 메모리로 읽어들여 빠른읽기를 수행할수 있지만 큰 파일을 읽을때는 비추
	f4 := "f4.txt"
	err = ioutil.WriteFile(f4, s, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 파일쓰기 방법5 : io.WriteString() 함수 사용
	f5, err := os.Create("f5.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = io.WriteString(f5, string(s))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}

// go run save.go
// ls -l f?.txt
// cat f?.txt
