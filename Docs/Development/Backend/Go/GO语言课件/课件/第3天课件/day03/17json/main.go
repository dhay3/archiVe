package main

import (
	"encoding/json"
	"fmt"
)

//json序列化

//Student 学生
// type Student struct {
// 	ID     int
// 	Gender string
// 	Name   string
// }

// Student 是一个结构体
//定义元信息：json tag
type student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
}

func main() {
	var stu1 = student{
		ID:     1,
		Gender: "男",
		Name:   "豪杰",
	}
	// 序列化：把编程语言里面的数据转换成 JSON 格式的字符串
	v, err := json.Marshal(stu1)
	if err != nil {
		fmt.Println("JSON格式化出错啦！")
		fmt.Println(err)
	}
	fmt.Println(v)               //[]byte
	fmt.Printf("%#v", string(v)) //把[]byte转换成string

	// str := "{\"ID\":1,\"Gender\":\"男\",\"Name\":\"豪杰\"}"
	// 	//反序列化：把满足JSON格式的字符串转换成 当前编程语言里面的对象
	// 	var stu2 = &Student{}
	// 	json.Unmarshal([]byte(str), stu2)
	// 	fmt.Println(stu2)
	// 	fmt.Printf("%T\n", stu2)
}
