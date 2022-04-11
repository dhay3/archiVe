package main

import "fmt"

// 没有参数没有返回值的函数
func sayHello() {
	fmt.Println("Hello 沙河！")
}

// 带参数的函数
func sayHi(name string) {
	fmt.Printf("Hello %s\n", name)
}

// 带参数和返回值
func add(a int, b int) int {
	ret := a + b
	return ret
}

func add2(a, b int) int {
	ret := a + b
	return ret
}

func add3(a, b int) (ret int) {
	ret = a + b
	return
}

// 参数类型可合并
// 命名返回值， return后面不写内容
// 没有命名返回值

// 可变参数
func add4(a int, b ...int) int {
	ret := a
	fmt.Println(a)
	fmt.Printf("b=%v type:%T\n", b, b)
	for _, v := range b {
		ret = ret + v
	}
	return ret
}

func add5() (int, int, int) {
	return 1, 2, 3
}
func main() {
	// sayHello()
	// sayHi("豪杰")
	// sayHi("nazha")
	// i := add(1, 2)
	// fmt.Println(i)
	fmt.Println(add4(1))
	fmt.Println(add4(1, 2))
	fmt.Println(add4(1, 2, 3))
}
