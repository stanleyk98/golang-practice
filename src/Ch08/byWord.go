package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func wordByWord(file string) error {
	var err error
	f, err := os.Open(file)

	if err != nil {
		return err
	}

	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("파일을 읽는도중 에러발생 %s", err)
			return err
		}
		// 입력된 라인을 단어단위로 구분짓는 정규표현식 사용(공백문자를 기준으로 분리)
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)

		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byword <file1> [<file2> ...]\n")
		return
	}

	for _, file := range flag.Args() {
		err := wordByWord(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
