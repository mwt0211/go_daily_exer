package main

import (
	"fmt"
	"os"
)

//文件写入
func main() {
	write()
}
func write() {
	fileWrite, err := os.OpenFile("file/write.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("err:%v", err)
	}
	var a string
	fmt.Println("请输入要追加的数据")
	for {
		scanf, _ := fmt.Scanf("%v\n", &a)
		if scanf == 0 {
			break
		}
	}

	fileWrite.Write([]byte(a))
}
