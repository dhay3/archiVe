package main

import "fmt"

//结构体
// 创在新的类型要使用type关键字
type student struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var haojie = student{
		name:   "豪杰",
		age:    19,
		gender: "男",
		hobby:  []string{"篮球", "足球", "双色球"},
	}

	//结构体支持.访问属性
	fmt.Println(haojie)
	fmt.Println(haojie.name)
	fmt.Println(haojie.age)
	fmt.Println(haojie.gender)
	fmt.Println(haojie.hobby)

	// 实例化方法1
	// struct是值类型的
	// 如果初始化时没有给属性（字段）设置对应的初始值，那么对应属性就是其类型的默认值
	var wangzhan = student{}
	fmt.Println(wangzhan.name)
	fmt.Println(wangzhan.age)
	fmt.Println(wangzhan.gender)
	fmt.Println(wangzhan.hobby)

	// 实例化方法2 new(T) T:表示类型或结构体
	var yawei = new(student)
	fmt.Println(yawei)
	// (*yawei).name
	yawei.name = "亚伟"
	yawei.age = 18
	fmt.Println(yawei.name, yawei.age)
	// 实例化方法3
	var nazha = &student{}
	fmt.Println(nazha)
	nazha.name = "沙河娜扎"
	fmt.Println(nazha.name)

	//结构体初始化
	//只填值初始化
	var stu1 = student{
		"豪杰",
		18,
		"男",
		[]string{"男人", "女人"},
	}
	fmt.Println(stu1.name, stu1.age)
	//键值对初始化
	var stu2 = &student{
		name:   "豪杰",
		gender: "男",
	}
	fmt.Println(stu2.name, stu2.age, stu2.gender)
}
