package main

import (
	"flag"
	"fmt"
	"strings"
)

// 이름 문자배열을 받는 구조체 정의
type NamesFlag struct {
	Names []string
}

// 이름을 받기
func (s *NamesFlag) GetNames() []string {
	return s.Names
}

// 받은 이름을 문자열 쪼개기
func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

// flag명은 중복호출 하지말기
func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("1회이상 플래그명을 사용할수 없음")
	}

	// 인수에 콤마를 기준으로 쪼개서 names 로 담기
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NamesFlag

	// 커맨드라인 들어오는 인수에 에러가 있을때 설명표시
	minusK := flag.Int("k", 0, "정수형")
	minus0 := flag.String("0", "TaeUk", "이름")
	flag.Var(&manyNames, "names", "콤마로 구분된 나열된 이름")

	flag.Parse()
	fmt.Println("-k: ", *minusK)
	fmt.Println("-0: ", *minus0)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("남은 커맨드라인 인수 : ")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}

// go run funwithflag.go -names=stanley,jin,won 1 two four
// go run funwithflag.go -id=stanley 1 tow four
// go run funwithflag.go -names=stanley -names=Stanley
