package main

import "fmt"

func main() {
	a := [...]int{1}       //数组 值类型
	b := a[:]              //引用类型
	fmt.Printf("%p\n", &a) //去找a这个变量对应的值的内存地址
	fmt.Printf("%p\n", b)  //取b指向的内存地址
	b[0] = 100
	fmt.Println(a, b)
	// 扩容策略
	fmt.Println("b的容量：", cap(b))
	b = append(b, 3, 4, 5, 6, 7, 8)
	fmt.Println("b的容量：", cap(b))
	b = append(b, 8)
	fmt.Println("b的容量：", cap(b))
	// 切片的总容易已经超过1024,再追加就0.25倍追加

	// make:用来给引用类型做初始化（申请内存空间）
	// new：用来创建值类型
	var s1 []int64 // 切片声明了没有初始化是不能使用
	// s1[0] = 100  // 不能这么做
	// s1 = append(s1, 100) // 可以这么用
	fmt.Println(s1)
	// 使用make初始化,切片扩容要申请内存空间
	s1 = make([]int64, 3)
	fmt.Println(s1, len(s1), cap(s1))
	s2 := make([]string, 3, 10)
	fmt.Println(s2, len(s2), cap(s2))
}
