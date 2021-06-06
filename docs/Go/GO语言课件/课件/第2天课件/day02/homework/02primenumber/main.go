package main

import "fmt"

// 打印200~1000之间的质数(素数)

func main() {
	for i := 200; i < 1000; i++ {
		// fmt.Println(i)
		flag := true
		//判断i是否为质数,如果是就打印,如果不是就不打印
		for j := 2; j < i; j++ {
			if i%j == 0 {
				//不是质数
				flag = false
				break
			}
		}
		//整个第二层的for循环结束了还没有break说明是质数
		if flag {
			fmt.Printf("%d是质数.", i)
		}
	}
}
