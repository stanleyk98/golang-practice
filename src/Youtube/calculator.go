package main

import (
	"bufio" // 한줄입력 패키지
	"fmt"
	"os"      // 인수 입력 패키지
	"strconv" //문자열 형변환 패키지
	"strings" //문자열 패키지
)

func main() {

	fmt.Println("숫자를 입력하세요~")
	reader := bufio.NewReader(os.Stdin) // 표준입력을 받아와서 reader 변수에 넣어라

	line, _ := reader.ReadString('\n') // 한줄을 읽어와서 line 변수에 넣어라
	line = strings.TrimSpace(line)     // 입력된 문자열에서 공백을 지운다 TrimSpace

	n1, _ := strconv.Atoi(line) // 문자열을 숫자로 형변환해서 n1 에 넣어라

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	n2, _ := strconv.Atoi(line)

	fmt.Printf("입력한 숫자는 %d,  %d 입니다  \n", n1, n2) // %d 는 정수로 표시하라는 뜻

	fmt.Println("연산자를 입력하세요~")

	line, _ = reader.ReadString('\n')
	line = strings.TrimSpace(line)

	// --- if else if 로 구현해보기 --- //

	// if line == "+" {
	// 	fmt.Printf("%d + %d = %d  \n", n1, n2, n1+n2)
	// } else if line == "-" {
	// 	fmt.Printf("%d - %d = %d  \n", n1, n2, n1-n2)
	// } else if line == "*" {
	// 	fmt.Printf("%d * %d = %d  \n", n1, n2, n1*n2)
	// } else if line == "/" {
	// 	fmt.Printf("%d / %d = %d  \n", n1, n2, n1/n2)
	// } else {
	// 	fmt.Println("연산자를 잘못 입력했습니다")
	// }

	// --- Switch Case 로 구현해보기 --- //

	switch line {
	case "+":
		fmt.Printf("%d + %d = %d  \n", n1, n2, n1+n2)
	case "-":
		fmt.Printf("%d - %d = %d  \n", n1, n2, n1-n2)
	case "*":
		fmt.Printf("%d * %d = %d  \n", n1, n2, n1*n2)
	case "/":
		fmt.Printf("%d / %d = %d  \n", n1, n2, n1/n2)
	default:
		fmt.Println("연산자를 잘못 입력했습니다")
	}

}
