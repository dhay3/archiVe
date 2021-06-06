package main

import "fmt"

// 函数可以作为变量、参数、返回值

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// calc是一个函数，它接收三个参数，返回一个int类型的返回值
// 其中，参数a和b是int类型
// 参数f 是一个函数类型，这个函数接收两个int类型的参数，返回一个int类型的返回值
func calc(a, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	f1 := add
	fmt.Printf("f1:%T\n", f1)
	fmt.Println(f1(10, 20))

	//把add当成一个参数传进calc中
	ret := calc(100, 200, add)
	fmt.Println(ret)

	// 把sub当成一个参数传进calc中
	ret = calc(100, 200, sub)
	fmt.Println(ret)
}
