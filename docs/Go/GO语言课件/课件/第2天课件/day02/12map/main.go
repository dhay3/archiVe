package main

import (
	"fmt"
	"sort"
)

func main() {
	// 声明map类型
	var m1 map[string]int
	// 未初始化就是 nil
	// 使用make初始化map
	m1 = make(map[string]int, 100)
	m1["nazha"] = 90
	m1["haojie"] = 100
	fmt.Println(m1)

	// 声明map类型时直接初始化
	m2 := map[string]string{
		"haojie": "hehe",
		"yawei":  "heihei",
	}
	fmt.Println(m2)

	// 判断map中是否存在某个键值对
	v, ok := m2["haojie"]
	// ok返回的是布尔值，能取到就返回true，取不到就返回false
	if !ok {
		fmt.Println("查无此人")
	} else {
		fmt.Println(v)
	}

	// map的遍历
	// 遍历键值对
	for k, v := range m2 {
		fmt.Println(k, v)
	}
	// 遍历键
	for k := range m2 {
		fmt.Println(k)
	}

	// 删除键值对
	delete(m2, "haojie")
	fmt.Println(m2)

	m3 := map[string]int{
		"haojie":   100,
		"nazha":    80,
		"cuicui":   70,
		"wangxin":  200,
		"qiaojing": 180,
	}
	// 将m2按照key的ASCII顺序打印键值对
	// 1. 先把key取出来放到切片里面
	fmt.Println(m3)
	var keys = make([]string, 0, 10)
	for key := range m3 {
		keys = append(keys, key)
	}
	fmt.Println(keys)
	// 2. 对key的切片做排序
	sort.Strings(keys)
	fmt.Println(keys)
	// 3. 按照排序后的key依次去map取值
	for _, key := range keys {
		fmt.Println(key, m3[key])
	}
}
