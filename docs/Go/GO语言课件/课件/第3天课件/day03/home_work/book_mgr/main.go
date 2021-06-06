package main

import (
	"fmt"
	"os"
)

// 使用函数实现一个简单的图书管理系统。
// 每本书有书名、作者、价格、上架信息，
// 用户可以在控制台添加书籍、修改书籍信息、打印所有的书籍列表。

// Book is a struct
type Book struct {
	title   string
	author  string
	price   float32
	publish bool
}

var (
	// AllBooks is a slice of Book pointer
	AllBooks []*Book
)

// 主菜单
func showMenu() {
	fmt.Println("1. 添加书籍")
	fmt.Println("2. 修改书籍信息")
	fmt.Println("3. 展示所有书籍")
	fmt.Println("4. 退出")
	fmt.Println()
}

func main() {
	for {
		showMenu()
		var input int
		fmt.Scanf("%d\n", &input) //从终端获取输入的数字
		switch input {
		case 1:
			AddBook()
		case 2:
			ModifyBook()
		case 3:
			showAllBook()
		case 4:
			os.Exit(0)
		}
	}
}
