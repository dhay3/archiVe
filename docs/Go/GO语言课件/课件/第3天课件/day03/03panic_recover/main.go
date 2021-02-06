package main

import "fmt"

func f1() {
	defer func() {
		// recover
		err := recover() //尝试将函数从当前的异常状态恢复过来
		fmt.Println("recover抓到了panic异常", err)
	}()
	var a []int
	a[0] = 100 //panic
	fmt.Println("panic之后")
}

// panic错误
func main() {
	f1()
	fmt.Println("这是main函数")
}
