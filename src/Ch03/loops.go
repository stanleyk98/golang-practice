package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {
		if i%20 == 0 { // i%20 : i 를 20으로 나눴을때 나머지 값 (즉, 20배수값)
			continue // 20배수일때 if 문 밖으로 나가지 않고 다음 i값 바로 진행되므로 print 함수 실행건너뛰게 됨
		}
		if i == 95 {
			break // i가 95가 되면 for 문 바로 종료 (if문 종료 + for문 종료) 되므로 print 함수 실행건너뛰게 됨
		}
		fmt.Print(i, " ")
	}
	fmt.Println() // 줄바꿈

	i := 10
	for {
		if i < 0 {
			break // i가 음수가 되면 for문 바로 종료(if문 종료 + for문 종료) 되므로 print 함수 실행건너뛰게 됨
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println() // 줄바꿈

	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression { // ok 와 anExpression 이 둘다 true 라면 계속 반복
		if i > 10 { // i가 11이 되면 anExpression 이 false 가 되고 다음 for문 돌때 ok = true, anExpression = false 가 되므로 for문 종료됨
			anExpression = false
		}
		fmt.Println(i, " ")
		i++
	}
	fmt.Println() // 줄바꿈

	anArray := [5]int{0, 1, -1, 2, -2} // int형으로 5개 인수의 배열 생성
	for i, value := range anArray {
		fmt.Println("index:", i, "value: ", value) // index 와 value 출력
	}
}
