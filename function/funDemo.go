package main

import "fmt"

func main() {
	//initSum(5, 8)


	intSum1()
	intSum1(22)
	intSum1(25, 80, 50, 100)
}
func initSum(x int, y int) int {
	sum := x + y
	fmt.Printf("%d + %d = %d\n", x, y, sum)
	return sum
}

//接受可变参数,在参数后加...表示参数的数量是可变的
//可变参数在函数体中是切片类型
func intSum1(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	//fmt.Printf("%T\n", x) //[]int
	fmt.Println(sum)
	return sum

}
