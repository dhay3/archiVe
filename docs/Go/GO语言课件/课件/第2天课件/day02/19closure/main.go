package main

import (
	"fmt"
	"strings"
)

// 闭包

// 定义一个函数f1
// f1不接受参数
// f1返回一个函数类型，这个函数不接收参数也没有返回值
func f1(num int) func(int) int {

	f := func(x int) int {
		fmt.Println("找到外层函数的变量num", num)
		return num + x
	}
	return f
}

// 闭包示例2
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

// 闭包示例3
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i // 引用了外层的base变量同时还修改了base
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	// 	ret := f1(100) // ret是一个闭包（闭包中引用了10）
	// 	fmt.Printf("%T\n", ret)
	// 	fmt.Println(ret(10))
	// 	fmt.Println(ret(20))
	// 	fmt.Println(ret(30))
	// 	fmt.Println(ret(40))

	// 	aviFunc := makeSuffixFunc(".avi")
	// 	fmt.Println(aviFunc("豪杰春香"))
	// 	textFunc := makeSuffixFunc(".txt")
	// 	fmt.Println(textFunc("豪杰春香"))

	f1, f2 := calc(100)
	fmt.Println(f1(150), f2(50))
	// 1. f1(150)  base += 150  --> base = 250
	// 2. f2(50)   base -= 50   --> base = 200
}
