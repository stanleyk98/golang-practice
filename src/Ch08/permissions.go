package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args

	// 커맨드라인에 파일명 기재안했을때 체크하기
	if len(arguments) == 1 {
		fmt.Printf("권한 체크할 파일명을 입력하세요~! \n")
		return
	}

	filename := arguments[1]

	// os.Stat() 함수를 호출하면 파일정보 구조체 반환
	// func (f *File) Stat() (FileInfo, error)
	info, _ := os.Stat(filename)

	// 파일정보 중 권한부분 확인 Mode함수
	mode := info.Mode()

	fmt.Println(filename, " <-- 이 파일의 권한은 ", mode.String()[1:10])
}

//go run permissions.go flo_cus_20201230.log
