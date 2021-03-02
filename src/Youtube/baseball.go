// 숫자야구 게임
// 두사람이 0~9까지 중에서 겹치지않는숫자 3개씩 정한다
// 상대방이 콜한 3자리 숫자가 자신의 숫자와 자리 위치 + 숫자일치 여부를 가린다
// 숫자는 일치하지만 자리가 다르면 볼
// 숫자와 자리까지 일치하면 스트라이크

/* Flow 분기표
1. 무작위 숫자 3개를 정한다
2. 사용자 입력한다
3. 비교한다
4. 3스트라이크면 "승리" 를 출력
5. 아니면 점수 알려주고 다시 입력받기
*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//seed 를 지정해둬야 랜덤다운 랜덤값생성됨
	rand.Seed(time.Now().UnixNano())

	// 무작위 숫자 3개를 생성한다
	numbers := MakeNumbers()

	cnt := 0

	for {
		cnt++

		// 사용자의 입력을 받는다
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3스트라이크인가?
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}
	// 게임끝이고 몇번만에 맞췄는지 출력
	fmt.Printf("%d 번만에 맞췄어요~ \n", cnt)
}

func MakeNumbers() [3]int {
	//0~9 사이에 겹치지않는 숫자 3개를 반환한다
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)  // 난수생성
			duplicated := false // Flag 변수 생성
			for j := 0; j < i; j++ {
				if rst[j] == n {
					//겹치는 숫자는 다시 뽑는다
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}

	fmt.Println(rst)
	return rst
}

func InputNumbers() [3]int {
	// 0~9 사이에 겹치지않는 숫재 3개를 입력받아 반환한다
	var rst [3]int

	for {
		var no int
		_, err := fmt.Scanf("%d", &no)
		if err != nil {
			fmt.Println("입력을 다시하세요")
			continue
		}

		idx := 0
		for no > 0 {
			// 1의 자리수, 10의 자리수, 100의 자리수를 구한다
			// 입력받는 값을 10으로 나눈 나머지 값이 1의 자리수
			n := no % 10
			no = no / 10
			rst[idx] = n
			idx++
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	return rst
}

func CompareNumbers(numbers, inputNumbers [3]int) bool {
	// 두개의 숫자 3개를 비교해서 결과를 반환한다
	return true
}

func PrintResult(result bool) {
	fmt.Println(result)
}

func IsGameEnd(result bool) bool {
	// 비교결과가 3 스트라이크인지 확인
	return result
}
