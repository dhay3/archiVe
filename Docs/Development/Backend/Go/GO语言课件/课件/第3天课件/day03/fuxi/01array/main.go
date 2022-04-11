package main

import "fmt"

func main() {
	// var a = [3][2]int{
	// 	{1, 2},
	// 	{3, 4},
	// 	{5, 6},
	// }
	// fmt.Println(a)
	b := [...]string{"beijing", "shanghai", "shenzhen"}
	// fmt.Println(b[1])
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}

	for index, value := range b {
		fmt.Println(index, value)
	}

	for index := range b {
		fmt.Println(index)
	}
	for _, value := range b {
		fmt.Println(value)
	}

}
