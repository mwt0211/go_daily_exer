package main

import "fmt"

func main() {
	//切片的copy
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)
	copy(b, a)
	fmt.Println(b)
	a[0] = 99
	fmt.Println(b) //[1 2 3 4 5]
	fmt.Println(a) //[99 2 3 4 5]
	fmt.Println("******************")
	c := b
	c[2] = 100
	fmt.Println(c) //[1 2 100 4 5]
	fmt.Println(b) //[1 2 100 4 5]
	//删除切片元素    append(Slice[0:index], Slice[index+1:]...)
	//eg:删掉广州
	citySlice := []string{"北京", "上海", "广州", "深圳"}
	for i := 0; i < len(citySlice); i++ {
		if citySlice[i] == "广州" {
			citySlice = append(citySlice[0:i], citySlice[i+1:]...)
		}
	}
	fmt.Println(citySlice)
}
