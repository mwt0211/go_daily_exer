package main

import (
	"fmt"
	"time"
)

//time
//每隔两秒打印一下当前的时间,20秒后结束打印
func main() {
	count := 0
	for {
		//time1:=time.Now().Format("2006-01-02 15:04:05")
		time.Sleep(time.Second * 2)
		//t1:=time.Now().Truncate(time.Second*3)
		//fmt.Printf("%s\n",t1.Format("2006-01-02 15:04:05"))
		t1 := time.Tick(time.Second * 3)
		fmt.Printf("%s\n", t1)
		count++
		//fmt.Printf("当前时间为:%s,第%d次打印\n",time1,count)
		if count > 10 {
			break
		}
	}

}
