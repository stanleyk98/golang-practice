package main

import (
	"flag"
	"fmt"
	"strings"
)

type NameFlag struct {
	Names []string
}

func (s *NameFlag) GetNames() []string {
	return s.Names
}

func (s *NameFlag) String() []string {
	return fmt.Sprint(s.Names)
}

func (s *NameFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("1회이상 플래그명을 사용할수 없음")
	}

	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}

func main() {
	var manyNames NameFlag

	minusK := flag.Int("k", 0, "정수형")
	minus0 := flag.String("0", "Stanley", "이름")

	flag.Var(&manyNames, "names", "콤마로 구분된 리스트")

	flag.Parse()
	fmt.Println("-k: ", *minusK)
	fmt.Println("-0: ", *minus0)

	for i, item := range manyNames.Names.GetNames() {
		fmt.Println(i, item)
	}

	fmt.Println("남은 커맨드라인 인수 : ")
	for index, val := range flag.Args() {
		fmt.Println(index, ":", val)
	}
}
