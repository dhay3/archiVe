package main

import "fmt"

func main() {
	// var a int
	// fmt.Println(a)
	// b := &a // 取变量a的内存地址
	// fmt.Printf("b=%v\n", b)
	// fmt.Printf("type b :%T\n", b)
	// c := "豪杰"
	// fmt.Printf("&c=%v\n", &c)
	// // b = &c //取c的内存地址不能赋值给b

	// d := 100
	// b = &d
	// fmt.Println(b)
	// // *取地址对应的值
	// fmt.Println(*b)
	// // 指针可以做逻辑判断
	// fmt.Println(b == &d)

	//指针的应用
	a := [3]int{1, 2, 3}
	modifyArray(a) //在函数中复制了数组赋值给了内部的a1
	fmt.Println(a)
	modifyArray2(&a)
	fmt.Println(a)
}

// 定义一个修改数组第一个元素为100的函数
func modifyArray(a1 [3]int) {
	a1[0] = 100 //只是修改的内部的a1这个数组
}

// 定义一个修改数组第一个元素为100的函数
// 接收的参数是一个数组的指针
func modifyArray2(a1 *[3]int) {
	// (*a1)[0] = 100 //只是修改的内部的a1这个数组
	//语法糖：因为Go语言中指针不支持修改
	a1[0] = 100 //只是修改的内部的a1这个数组
}
