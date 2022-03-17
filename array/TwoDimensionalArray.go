package main

import "fmt"

func main() {
	//二维数组
	city := [3][2]string{{"广东", "上海"}, {"浙江", "河南"}, {"南京", "北京"}}
	fmt.Println(city)
	//fmt.Println(city[1][1])
	//遍历二维数组
	for key, v1 := range city {
		fmt.Println(key, v1)
		for _, value := range v1 {
			fmt.Println(value)
		}

	}
}
