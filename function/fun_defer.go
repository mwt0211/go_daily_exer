package main

import "fmt"

var globalVariable string = "全局变量"

func main() {
	//defer延迟调用
	fmt.Println("start")
	defer fmt.Println("defer1")
	defer fmt.Println("defer2")
	fmt.Println("middle")
	defer fmt.Println("defer3")
	fmt.Println("end")

}
