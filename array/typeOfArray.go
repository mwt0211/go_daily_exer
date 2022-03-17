package main

import "fmt"

func main() {
	//数组是值类型
	var intArray = [3]int{1, 2, 3}
	var stArray = [3][2]string{{"北京", "上海"}, {"重庆", "南京"}, {"浙江", "河南"}}
	fmt.Println(intArray)

	f1(intArray)
	fmt.Println(intArray)
	fmt.Println("-----------------")
	fmt.Println(stArray)
	f2(stArray)
	fmt.Println(stArray)
}
func f1(a [3]int) {
	a[0] = 10
}
func f2(b [3][2]string) {
	b[1][1] = "小贾"
}
