package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}

func main() { //开启一个主goroutine去执行main函数
	/*//runtime.GOMAXPROCS(1):先执行一个goroutine,执行完成之后才会执行另一个
	runtime.GOMAXPROCS(1)//只占用一个CPU核心*/
	runtime.GOMAXPROCS(2) //占用两个CPU核心并行执行任务,不保证其顺序性
	wg.Add(2)             //相当于定义有几个goroutine
	go speak()
	go hello() //开启一个goroutine去执行函数
	wg.Wait()  //等所有子goroutine全部执行完再释放
	//具体的执行顺序是无序的

}
func hello() {
	for i := 0; i < 10; i++ {
		fmt.Println("hello 娜扎", i)
	}
	wg.Done() //告知主goroutine函数执行完毕
}
func speak() {
	for i := 0; i < 10; i++ {
		fmt.Println("说话", i)
	}
	wg.Done()
}
