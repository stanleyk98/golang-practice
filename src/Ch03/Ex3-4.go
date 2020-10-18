// 연습문제 3-4
// parse.Date.go 파일을 수정해보기

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myDate string
	if len(os.Args) != 2 {
		fmt.Println("다음 형식으로 입력해주세요 DD-Jan-YYYY")
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myDate = os.Args[1]

	d, err := time.Parse("02-January-2020", myDate)

	if err == nil {
		fmt.Println("Full: ", d)
		//fmt.Println("Time: ", d.Day(), d.Month(), d.Year())
		fmt.Println("Time: ", d.Month(), "월", d.Day(), "일", d.Year(), "년")
	} else {
		fmt.Println(err)
	}
}
