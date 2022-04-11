package main

import "fmt"

//全局变量和局部变量
var globalVariable = "陈小瑞"

func main() {
	TestGlobal("你好啊")
	TestLocal("我好像喜欢上你了")

}
func TestGlobal(global string) {
	fmt.Println(global + globalVariable)
}
func TestLocal(local string) string {
	globalVariable = "瑞瑞"
	ret := local + globalVariable

	fmt.Println(ret)
	return ret
}
