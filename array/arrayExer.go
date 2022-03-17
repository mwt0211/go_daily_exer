package main

import "fmt"

func main() {
	//1.求数组的和
	var sumArray = [5]int{1, 3, 5, 7, 8}
	var sum = 0
	for _, sumArray := range sumArray {
		sum += sumArray
	}
	fmt.Printf("sumArray的和为:%d\n", sum)

	//2.找出数组中和为指定值的两个元素的下标
	//例如:数组{1,3,5,7,8}找出和为8的两个元素的下标
	var index = [5]int{1, 3, 5, 7, 8}
	for i := 0; i < len(index); i++ {
		for j := i; j < len(index); j++ {
			if index[i]+index[j] == 8 {
				fmt.Printf("索引为(%d,%d),", i, j)
			}
		}

	}
}
