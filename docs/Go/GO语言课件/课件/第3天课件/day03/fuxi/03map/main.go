package main

import "fmt"

//map
func main() {
	// 光声明map但是没有初始化 它就是nil
	// var m1 map[string]int
	// m1 = make(map[string]int, 10)

	m1 := make(map[string]int, 10)
	m1["沙河"] = 100
	fmt.Println(m1)
	// map中取 key的值 一种固定写法
	v, ok := m1["西二旗"]
	if ok {
		fmt.Println("m1中有西二旗这个key")
		fmt.Println(v)
	} else {
		fmt.Println("m1中没有西二旗这个key")
	}
}
