package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() { //开启一个主goroutine去执行main函数

	for i := 0; i < 1000; i++ {
		//匿名函数,不含参,变量值为外部的
		/*		//会出现多个相同的打印输出语句,其原因为其匿名函数也为闭包,其变量I为外部作用域的,并不是自己内部的,需要从外部找.具有延时性
				go func() {
				wg.Add(1)
					fmt.Println("hello", i)
					wg.Done()
				}()*/
		go func(i int) {
			wg.Add(1)
			//此刻的i为入参的I,并不是外部传入的,所以能保证其不一致性
			fmt.Println("hello", i)
			wg.Done()
		}(i) //此处的i为for循环中的实时性i,然后将此处的I作为入参传给匿名函数,保证其实时性
	}
	fmt.Println("hello main")
	wg.Wait()
}
