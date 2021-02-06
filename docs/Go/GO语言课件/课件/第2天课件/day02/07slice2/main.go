package main

import "fmt"

func main() {
	var a = []int{} //空切片
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1) // [1 1 1]  3  4
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
	a = append(a, 1)
	fmt.Printf("a:%v len:%d  cap:%d ptr:%p\n", a, len(a), cap(a), a)
}
