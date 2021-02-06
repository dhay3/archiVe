package main

import "fmt"

//结构体的内存布局
// 内存是以字节为单位的十六进制数
// 1字节 = 8位 = 8bit

func main() {
	type test struct {
		a int16
		b int16
		c int16
	}

	var t = test{
		a: 1,
		b: 2,
		c: 3,
	}
	fmt.Println(&(t.a))
	fmt.Println(&(t.b))
	fmt.Println(&(t.c))
}
