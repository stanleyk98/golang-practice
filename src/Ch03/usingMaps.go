package main

import (
	"fmt"
)

func main() {
	iMap := make(map[string]int)
	iMap["empno"] = 150009
	iMap["age"] = 300
	fmt.Println("iMap: ", iMap)

	anotherMap := map[string]int{
		"사번": 150009,
		"나이": 300,
	}

	fmt.Println("anotherMap: ", anotherMap)
	delete(anotherMap, "사번")
	delete(anotherMap, "사번")
	delete(anotherMap, "사번")
	fmt.Println("anotherMap: ", anotherMap)

	// 맵에 해당 값이 있는지 확인
	if empno, ok := iMap["148800"]; ok {
		fmt.Println(empno, ok)
	} else {
		fmt.Println("값없음")
	}

	// _, ok := iMap["doesItExist"]
	// if ok {
	// 	fmt.Println("Exists~!")
	// } else {
	// 	fmt.Println("Does NOT Exists")
	// }
	// for key, value := range iMap {
	// 	fmt.Println(key, value)
	// }

}
