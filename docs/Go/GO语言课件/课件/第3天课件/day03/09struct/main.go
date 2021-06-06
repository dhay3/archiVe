package main

import "fmt"

type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

//自己实现一个构造函数
func newStudent(n string, age int, g string, h []string) *student {
	return &student{
		name:   n,
		age:    age,
		gender: g,
		hobby:  h,
	}
}

func main() {
	hobbySlice := []string{"篮球", "球"}
	haojie := newStudent("豪杰", 18, "男", hobbySlice)
	fmt.Println(haojie.name)
}
