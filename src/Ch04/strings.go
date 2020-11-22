package main

import (
	"fmt"
)

func main() {
	const sLiteral = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	fmt.Println(sLiteral)
	fmt.Printf("x: %x\n", sLiteral)
	fmt.Printf("sLiteral 의 길이: %d\n", len(sLiteral))

	for i := 0; i < len(sLiteral); i++ {
		fmt.Printf("%x ", sLiteral[i])
	}
	fmt.Println()

	fmt.Printf("q: %q\n", sLiteral)
	fmt.Printf("+q: %+q\n", sLiteral)
	fmt.Printf(" x: % x\n", sLiteral)

	fmt.Printf("s: As string: %s\n", sLiteral)

	s2 := ""
	for x, y := range s2 {
		fmt.Printf("%#U 문자하나씩 %d\n", y, x)
	}
	fmt.Printf("s2 길이: %d\n", len(s2))

	const s3 = "ab12AB"
	fmt.Println("s3:", s3)
	fmt.Printf("x: % x\n", s3)

	fmt.Printf("s3의 길이: %d\n", len(s3))

	for i := 0; i < len(s3); i++ {
		fmt.Printf("%x ", s3[i])
	}
	fmt.Println()
}
