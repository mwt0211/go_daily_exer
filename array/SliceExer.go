package main

import (
	"fmt"
	"sort"
)

func main() {
	//1.判断下面的输出结果
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	//2.使用内置的sort对数组var intArr=[5]int{6,9,2,5,4}排序
	var intArr = [5]int{6, 9, 2, 5, 4}
	sort.Ints(intArr[:])
	fmt.Println(intArr)
}
