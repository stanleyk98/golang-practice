package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("tmptext.txt")

	if err != nil {
		fmt.Println("파일 여는 중 에러 발생")
	}
	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanWords)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())

	}

	if err := fileScanner.Err(); err != nil {
		fmt.Println("파일 읽는 중 에러발생")
	}
	file.Close()

}

// go run ch08_ex12.go
// file 을 byWord.go 예제처럼 변수에 담을려면 어떻게?
