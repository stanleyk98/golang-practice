package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func lineByLine(file string) error {
	var err error

	// 에러체크
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// 읽을 파일을 열었으면 reader 를 새로 생성
	r := bufio.NewReader(f)
	for {
		// 줄바꿈 문자 \n 나올때까지 string 을 read 하고 라인 끝에서 break
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("파일읽는 도중 에러 발생 %s", err)
			break
		}

		fmt.Println(strings.TrimSpace(line))

	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage : byLine <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := lineByLine(file)

		if err != nil {
			fmt.Println(err)
		}
	}
}

// go run ch08_ex05.go tmptext.txt
// 왜 공백이 제거되지 않을까?
