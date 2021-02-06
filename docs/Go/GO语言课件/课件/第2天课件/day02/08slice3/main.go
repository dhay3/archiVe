package main

import "fmt"

// 切片是引用类型
func main() {
	a := []int{1, 2, 3}
	b := a //直接赋值
	// var c []int //还没有申请内存
	// c = make([]int, 3, 3)

	var c = []int{0, 0}

	num := copy(c, a) //深拷贝

	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(num)

}
