package main

import (
	"fmt"
	"strings"
)

//统计一个字符串中,每个单词出现的次数

func main() {
	//例如:"how do you do, I like what you like"
	var str = "how do you do , I like what you like"
	words := strings.Split(str, " ")
	fmt.Println(words)
	var strMap = make(map[string]int, 25)
	for _, word := range words {
		v, ok := strMap[word]
		if ok {
			strMap[word] = v + 1
		} else {
			strMap[word] = 1
		}
	}
	for key, count := range strMap {
		fmt.Println(key, count)
	}

}
