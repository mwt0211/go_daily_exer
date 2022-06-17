package main

import "fmt"

func main() {
	//initSum(5, 8)

	
	intSum1()
	intSum1(22)
	intSum1(25, 80, 50, 100)
	x,y:=calc(19, 5, 6, 8,9)
	fmt.Println(x,y)
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
func calc(a int, y ...int) (sum, sub int) {
	sum = a
	for _, v := range y {
		sum = sum + v
	}
	fmt.Printf("和为:%d\n", sum)
	sub = a
	for _, v := range y {
		sub = sub - v

	}
	fmt.Printf("差为:%d\n", sub)
	return sum, sub
}
