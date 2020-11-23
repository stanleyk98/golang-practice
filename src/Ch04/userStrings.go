package main

import (
	"fmt"
	s "strings"
	"unicode"
)

//fmt.Printf 함수를 변수에 넣어서 변수를 호출해서 간결히 사용할수도 있지만 복잡해지면 코드가독성 떨어짐
var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there~~!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THere"))
	f("%s\n", s.Title("tHis wiLL be A title!"))

	//EqualFold : string 두개를 같은지 비교한다
	f("EqualFold : %v\n", s.EqualFold("Stanley", "STAnley"))
	f("EqualFold : %v\n", s.EqualFold("Stanley", "stanle"))

	//Prefix : string 이 매개변수로 시작하면 true (대소문자 구분함)
	f("Prefix : %v\n", s.HasPrefix("Stanley", "St"))
	f("Prefix : %v\n", s.HasPrefix("Stanley", "st"))

	//Suffix : string 이 매개변수로 끝나면 true (대소문자 구분함)
	f("Suffix : %v\n", s.HasSuffix("Stanley", "ey"))
	f("Suffix : %v\n", s.HasSuffix("Stanley", "EY"))

	//Index :
	f("Index : %v\n", s.Index("Stanley", "st"))
	f("Index : %v\n", s.Index("Stanley", "St"))

	//Count : string 에서 지정한 매개변수를 중복되지않고 카운트
	f("Count : %v\n", s.Count("Stanley", "a"))
	f("Count : %v\n", s.Count("Stanley", "l"))

	//Repeat :
	f("Repeat : %s\n", s.Repeat("ab", 5))

	//TrimSpace :
	//TrimLeft :
	//TrimRight :
	f("TrimSpace : %s\n", s.TrimSpace(" \tHis is a line. \n"))
	f("TrimLeft : %s", s.TrimLeft(" \tHis is a\t line. \n", "\n\t "))
	f("TrimRight : %s\n", s.TrimRight(" \tHis is a\t line. \n", "\n\t "))

	//Compare : stirng 을 비교해서 같으면 0 , 다르면 -1 또는 1 을 리턴
	f("Compare : %v\n", s.Compare("Stanley", "STANLEY"))
	f("Compare : %v\n", s.Compare("Stanley", "Stanley"))
	f("Compare : %v\n", s.Compare("STANLEY", "STanley"))

	//Field : string 을 공백문자 기준으로 쪼갠다
	f("Field : %v\n", s.Fields("This is a string~!"))
	f("Field : %v\n", s.Fields("Thisis\na\tstring~!"))

	//Split : string 을 매개변수 기준으로 쪼갠다
	f("%s\n", s.Split("abcd efg", ""))

	//Replace : String 중에 세번째 매개변수로 교체하고 횟수지정가능 (-로 지정하면 무제한)
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%s\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join : %s\n", s.Join(lines, "+++"))

	f("SplitAfter : %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}

	f("TrimFunc : %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))

}
