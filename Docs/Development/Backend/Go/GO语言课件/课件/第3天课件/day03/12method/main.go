package main

import "fmt"

//函数是谁都可以调用的。
//方法就是某个具体的类型才能调用的函数
type people struct {
	name   string
	gender string
}

//函数指定接受者之后就是方法
// 在go语言中约定成俗不用this也不用self,而是使用后面类型的首字母的小写
func (p *people) dream() {
	p.gender = "男"
	fmt.Printf("%s的梦想是不用上班也有钱拿！\n", p.name)
}
func main() {
	var haojie = &people{
		name:   "豪杰",
		gender: "爷们",
	}
	// (&haojie).dream()
	haojie.dream()
	fmt.Println(haojie.gender)
}
