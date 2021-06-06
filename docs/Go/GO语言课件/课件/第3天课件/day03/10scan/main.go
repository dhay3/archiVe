package main

import "fmt"

//从终端获取用户的输入内容

func main() {
	var (
		name    string
		age     int
		married bool
	)
	fmt.Println(name, age, married)
	// fmt.Scan(&name, &age, &married)
	// fmt.Scanf("name:%s age:%d married:%t\n", &name, &age, &married)
	fmt.Scanln(&name, &age, &married)
	fmt.Println(name, age, married)

	// 1. 打印菜单
	// 2. 等待用户输入菜单选项
	// 3. 添加书籍的函数
	// 4. 修改书籍的函数
	// 5. 展示书籍的函数
	// 6. 退出 os.Exit(0)
}
