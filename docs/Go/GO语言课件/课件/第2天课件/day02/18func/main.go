package main

import "fmt"

// 匿名函数
func main() {
	func() {
		fmt.Println("Hello world!")
	}()
}
