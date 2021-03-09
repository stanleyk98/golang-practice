package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var buffer bytes.Buffer
	buffer.Write([]byte("Original string"))

	// Fprint : print 앞에 F 가 붙으면 파일 입출력에 해당된다고 볼것
	// buffer 라는 변수에 내용 추가
	fmt.Fprintf(&buffer, " Add string~!\n")

	// 표준출력으로 buffer 변수의 내용을 출력하기 위해 writeto() 함수 사용
	buffer.WriteTo(os.Stdout)
	// 이미 출력해서 buffer 변수는 clear 되어있으므로 출력데이터 없음 --> 왜?
	buffer.WriteTo(os.Stdout)

	// buffer 변수를 리셋 --> 비워져 있는데 리셋은 왜?
	buffer.Reset()
	buffer.Write([]byte("Mastering Go~"))

	r := bytes.NewReader([]byte(buffer.String()))

	// 이건 뭘 출력하는지?
	fmt.Println(buffer.String())

	for {
		// cap 3인 byte 슬라이스 생성
		b := make([]byte, 3)

		// 파일읽기
		n, err := r.Read(b)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("읽은 내용은 %s 이고 %d Bytes 이다.\n", b, n)
	}
}

// go run bytes.go
