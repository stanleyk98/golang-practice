package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		io.WriteString(os.Stdout, scanner.Text())
		io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	filename := ""
	arguments := os.Args

	// 커맨드라인 인수에 입력값 없으면 표준출력에 표준입력값을 복사한다
	if len(arguments) == 1 {
		io.Copy(os.Stdout, os.Stdin)
		return
	}

	for i := 1; i < len(arguments); i++ {
		filename = arguments[i]
		err := printFile(filename)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// go run cat.go   --> 실행하면 입력한 인수에 대해 출력함
// go run cat.go flo_cus_20201230.log   --> 파일 읽어와서 출력하는 것과 다른게 무엇?
