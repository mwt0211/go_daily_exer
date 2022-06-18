package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() { //开启一个主goroutine去执行main函数

	wg.Add(2) //相当于定义有几个goroutine
	go speak()
	go hello() //开启一个goroutine去执行函数
	fmt.Println("hello main")
	fmt.Println("***********")
	wg.Wait() //等所有子goroutine全部执行完再释放
	//具体的执行顺序是无序的

}
func hello() {

	fmt.Println("hello 娜扎")
	wg.Done() //告知主goroutine函数执行完毕
}
func speak() {
	fmt.Println("讲话")
	wg.Done()
}
