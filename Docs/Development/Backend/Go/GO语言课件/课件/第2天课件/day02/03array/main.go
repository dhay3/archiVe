package main

import "fmt"

func main() {
	//声明
	var a [5]int  //定义一个长度为5存放int类型的数组
	var b [10]int // 定义一个长度为10存放int类型的数组
	// 初始化
	a = [5]int{1, 2, 3, 4, 5}
	b = [10]int{1, 2, 3, 4}
	fmt.Println(a)
	fmt.Println(b)

	// var c [3]string = [3]string{"北京", "上海", "深圳"}
	// var c = [3]string{"北京", "上海", "深圳"}
	// fmt.Println(c)
	// // ...表示让编译器去数一下有多少个初始值，然后给变量赋值类型
	// var d = [...]int{1, 23, 3, 45545, 5642323, 3242, 1}
	// fmt.Println(d)
	// fmt.Printf("c:%T  d:%T\n", c, d)

	// // 根据索引值初始化
	// var e [20]int
	// e = [20]int{19: 1}
	// fmt.Println(e)

	// //数组的基本使用
	// fmt.Println(e[19])

	// 遍历数组的方式1
	// for i := 0; i < len(a); i++ {
	// 	fmt.Println(a[i])
	// }

	// for range循环
	for index, value := range a {
		fmt.Println(index, a[index], value)
	}
}
