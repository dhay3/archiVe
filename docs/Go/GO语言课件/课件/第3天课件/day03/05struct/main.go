package main

import "fmt"

// NewInt 是一个新的类型
type NewInt int

// 类型别名:只存在代码编写过程中，代码编译之后根本不存在haojie
// 提高代码的可读性
type haojie = int

// byte uint8

func main() {
	var a NewInt
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	var b haojie
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c byte
	fmt.Println(c)
	fmt.Printf("%T\n", c)
}
