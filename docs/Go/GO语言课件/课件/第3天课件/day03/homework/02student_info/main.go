package main

import "fmt"

// 设计一个程序，存储学员信息:id 姓名 年龄 分数
// 能根据id获取学员信息
func main() {
	studentMap := make(map[string]map[string]int, 10)
	// 初始化内层的map
	studentMap["豪杰"] = make(map[string]int, 3)
	studentMap["豪杰"]["id"] = 1
	studentMap["豪杰"]["age"] = 18
	studentMap["豪杰"]["score"] = 90

	studentMap["娜扎"] = make(map[string]int, 3)
	studentMap["娜扎"]["id"] = 2
	studentMap["娜扎"]["age"] = 28
	studentMap["娜扎"]["score"] = 100

	for k, v := range studentMap {
		fmt.Println(k)
		for k2, v2 := range v {
			fmt.Println(k2, v2)
		}
	}
	// 根据id获取学员信息
	for k, v := range studentMap {
		id, ok := v["id"]
		if ok {
			if id == 1 {
				fmt.Println("name", k)
				for k2, v2 := range v {
					fmt.Println(k2, v2)
				}
			}
		} else {
			fmt.Println("查无此人！")
		}
	}
}
