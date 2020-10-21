package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	var myTime string

	if len(os.Args) != 2 {
		fmt.Printf("usage: %s string\n", filepath.Base(os.Args[0]))
		return
	}
	myTime = os.Args[1]
	//Mon Jan 2 15:04:05 MST 2006

	d, err := time.Parse("15:04", myTime)
	if err == nil {
		fmt.Println("Full: ", d)
		fmt.Println("Time: ", d.Hour(), d.Minute())
		fmt.Println("확인", os.Args[0])
		fmt.Println("확인", os.Args[1])
	} else {
		fmt.Println(err)
	}
}
