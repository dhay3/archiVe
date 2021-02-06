package main

import "fmt"

// 切片中删除某个元素
func main() {
	a := []string{"北京", "上海", "深圳", "广州"}
	// a[:1] = ["北京"]
	// a[2:] = ["深圳" "广州"]
	// a = append(a, "成都")
	// a = append(a, ["深圳" "广州"]...)
	a = append(a[:1], a[2:]...)
	fmt.Println(a)
}
