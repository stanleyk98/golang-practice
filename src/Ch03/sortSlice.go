package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Noel", 200, 80})
	mySlice = append(mySlice, aStructure{"Jackson", 180, 60})
	mySlice = append(mySlice, aStructure{"Stanley", 300, 100})
	mySlice = append(mySlice, aStructure{"Terry", 175, 60})
	mySlice = append(mySlice, aStructure{"Edwin", 184, 50})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height // 키 작은순으로 정렬
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height // 키 큰순으로 정렬
	})
	fmt.Println(">:", mySlice)
}
