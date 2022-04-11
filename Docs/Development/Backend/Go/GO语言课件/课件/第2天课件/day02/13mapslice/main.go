package main

import "fmt"

func main() {
	// 1. 切片的元素是map
	// 初始化切片
	mapSlice := make([]map[string]int, 3, 10)
	// [map[aaa][10] map[bbb]100 ]
	fmt.Println(mapSlice)
	mapSlice = append(mapSlice, map[string]int{"aaa": 10})
	mapSlice = append(mapSlice, map[string]int{"bbb": 100})
	fmt.Println(mapSlice)
	// 此时mapSlice[2] 对应的map还没初始化
	mapSlice[2] = make(map[string]int, 10)
	mapSlice[2]["age"] = 18
	fmt.Println(mapSlice)

	// map中元素的值是切片
	// 对外层的map做初始化
	sliceMap := make(map[string][]int, 10)
	// 对map的值（切片）做初始化
	sliceMap["haojie"] = make([]int, 3, 10)
	sliceMap["haojie"][0] = 1
	sliceMap["haojie"][1] = 2
	sliceMap["haojie"][2] = 3
	fmt.Println(sliceMap)
}
