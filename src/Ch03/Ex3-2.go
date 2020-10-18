//연습문제 3-2
// parseTime.go 파일의 시간 변환을 임의 수정해보기

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string

	const shortForm = "2006-Jan-02"

	if len(os.Args) != 2 {
		fmt.Println("시간을 입력해주세요 HH:MM")
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myTime = os.Args[1]

	d, err := time.Parse("15:04", myTime)
	//d, err := time.Parse(shortForm, myTime)

	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), "시", d.Minute(), "분")

	} else {
		fmt.Println(err)
	}
	//t,_ = time.parse(shortForm, myTime)
}
