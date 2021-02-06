package main

import "fmt"

// 全局变量
var a = 10

func testGlobal() {
	// 函数中访问a
	// 1. 现在自己的作用域中找a
	// 2. 如果没有就找外层找，这个函数的外层就是全局
	a := 100 // 重新在函数中申请了一个局部变量a
	b := 200
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	testGlobal()
	fmt.Println(a)
	// fmt.Println(b)//内层可以往外层找，外层看不到内层的变量

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// fmt.Println(i)// 变量i只在for循环的语句块里生效
}
