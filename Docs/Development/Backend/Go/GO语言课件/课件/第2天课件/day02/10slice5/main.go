package main

import "fmt"

func main() {
	// 定义一个数组
	a := [...]int{1, 3, 5, 7, 9, 11, 13}

	// 基于数组得到一个切片
	b := a[:]
	// 修改切片中的第一个元素为100
	b[0] = 100
	// 打印数组中第一个元素的值
	fmt.Println(a[0])
	fmt.Printf("b:%p\n", b)

	c := a[2:5]
	fmt.Println(c)      //[5 7 9]
	fmt.Println(len(c)) //3
	fmt.Println(cap(c)) //5
	fmt.Printf("c:%p\n", c)

	d := c[:5]
	fmt.Println(d)
	fmt.Println(len(d)) //5
	fmt.Println(cap(d)) //5
	fmt.Printf("d:%p\n", d)

	e := d[2:]
	fmt.Println(e)      //[9 11 13]
	fmt.Println(len(e)) //3
	fmt.Println(cap(e)) //3
	fmt.Printf("e:%p\n", e)

	e = append(e, 100, 200)
	fmt.Println(e)
	fmt.Println(len(e)) //5
	fmt.Println(cap(e)) //6
	fmt.Printf("e:%p\n", e)

	e[0] = 900
	fmt.Println(d)
	fmt.Println(c)
	fmt.Println(e)

	m := []int{1, 2, 3, 4}
	fmt.Println(m)
	fmt.Printf("%p\n", m)
	m = append(m[:1], m[3:]...)
	fmt.Printf("%p\n cap:%d  len:%d\n", m, cap(m), len(m))
	fmt.Println(m)
}
