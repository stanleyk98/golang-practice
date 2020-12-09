package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	arguments := os.Args

	timeT, _ := time.Parse("2006-01-02", arguments[1])
	fmt.Println(timeT)

	// timeT, _ = time.Parse("06-01-02", "20-04-14")
	// fmt.Println(timeT)

	// timeT, _ = time.Parse("2006-Jan-02", "2020-Apr-14")
	// fmt.Println(timeT)

	// timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Apr-14 Tuesday 23:19:25")
	// fmt.Println(timeT)

	// timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Apr-14 Tuesday 11:19:25 PM IST +05:30")
	// fmt.Println(timeT)

}
