package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	//按照某个固定的顺序遍历map
	var scoreMap = make(map[string]int, 100)
	//添加50个键值对
	for i := 0; i < 50; i++ {
		key := fmt.Sprintf("stu%02d", i)
		value := rand.Intn(100) //0~99随机数
		scoreMap[key] = value
	}
	//for key, value := range scoreMap {
	//	fmt.Println(key,value)//随机打印
	//}
	//要求按照学号递增的顺序打印
	//1.取出所有的key
	keys := make([]string, 0, 100)
	for key, _ := range scoreMap {
		keys = append(keys, key)
	}
	//2.对key排序
	sort.Strings(keys)
	//3.按照key输出
	/*	for _, scorekey := range keys {
		fmt.Println(scorekey, scoreMap[scorekey])
	}*/

	//元素类型为map的Slice
	var mapSlice = make([]map[string]int, 8, 8) //完成对切片的初始化
	fmt.Println(mapSlice[0] == nil)             //true:说明尚未完成对map的初始化
	mapSlice[0] = make(map[string]int, 8)
	for i := 0; i < len(mapSlice); i++ {
		mapSlice[i] = make(map[string]int) //完成对map的初始化
	}
	mapSlice[5]["小红"] = 80
	fmt.Println(mapSlice)

	//值类型为切片的map
	var SliceMap = make(map[string][]int, 8)
	v, ok := SliceMap["小贾"]
	if ok {
		fmt.Println(SliceMap["小贾"])
	} else {
		SliceMap["小贾"] = make([]int, 8, 8) //完成对切片的初始化
		SliceMap["小贾"][0] = 18
		fmt.Println(v)
		fmt.Println(SliceMap)
	}
}
