package main

import "fmt"

// 1. 定义一个数组[1,3,5,7,8]，求数组中所有元素的和
func main() {
	a1 := [...]int{1, 3, 5, 7, 8}
	fmt.Println(a1)
	// // 1. 数组求和
	// sum := 0
	// // := 相当于声明变量并且赋值
	// for _, v := range a1 {
	// 	// fmt.Println(v)
	// 	sum = sum + v
	// }
	// fmt.Println(sum)

	// 2. 找出数组中和为指定值的两个元素的下标，比如从数组[1,3,5,7,8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
	//   遍历数组，
	// 2.1依次取出每个元素
	// 2.2 计算一下 other= 8-当前值
	// 2.3 在不在数组中，在的话把索引拿出来
	// [1, 3, 5, 7, 8]
	for index, value := range a1 {
		other := 9 - value
		for index2 := index + 1; index2 < len(a1); index2++ {
			if a1[index2] == other {
				// 另一半在数组中，把它们的索引打印出来
				fmt.Printf("它们的索引是：(%d %d)\n", index, index2)
			}
		}
	}
}
