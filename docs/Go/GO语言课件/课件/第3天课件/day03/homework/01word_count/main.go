package main

import (
	"fmt"
	"strings"
)

//写一个程序，
// 统计一个字符串中每个单词出现的次数。
// 比如："how do you do"中how=1 do=2 you=1。

func main() {
	s := "hello hello how are you i am fine thank you"
	// 1. 看一下字符串中都有哪些单词(用空格分割字符串得到字符串切片)
	wordSlice := strings.Split(s, " ")
	fmt.Println(wordSlice)
	// 2. 挨个数一数单词
	// 3. 找个适合的数据类型
	// 3.1 找个适合的数据类型
	wordMap := make(map[string]int, len(wordSlice))
	// 3.2 把结果存起来
	for _, word := range wordSlice {
		// fmt.Println(value)
		v, ok := wordMap[word]
		if ok {
			// 3.2.2 如果这个单词在map中就在原来的次数基础上+1
			wordMap[word] = v + 1
		} else {
			// 3.2.1 如果这个单词不在map中就添加键值对，让次数为1
			wordMap[word] = 1
		}
	}
	// 遍历map
	for k, v := range wordMap {
		fmt.Println(k, v)
	}
}
