package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func lineRead(log string) error {
	var err error

	f, err := os.Open(log)

	if err != nil {
		return err
	}
	defer f.Close()

	// 교재방법활용
	// r := bufio.NewReader(f)
	// for {
	// 	line, err := r.ReadString(',')
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		fmt.Printf("파일읽는동안 에러발생 %s", err)
	// 		break
	// 	}
	// 	fmt.Print(line)

	// }

	// 다른방법
	fileScanner := bufio.NewScanner(f)

	//한줄씩 읽기
	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
	}
	return nil
}

func main() {
	// csv 파일 생성
	file, err := os.Create("./output_test.csv")

	//에러나면 멘붕
	if err != nil {
		panic(err)
	}

	// csv writer 생성
	wr := csv.NewWriter(bufio.NewWriter(file))

	// csv 내용 쓰기
	// log 파일을 한줄씩 읽어와서 csv 파일에 한줄씩 입력

	lineRead("./flo_login_20201219.log")

	// 코드정리해야함....
	wr.Write([]string{"A", "0.25"})
	wr.Write([]string{"B", "55.70"})
	wr.Flush()
}
