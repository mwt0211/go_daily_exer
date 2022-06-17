package main

import "fmt"

//全局变量和局部变量
var globalVariable = "陈小瑞"

func main() {
	TestGlobal("你好啊")
	TestLocal("我好像喜欢上你了")
calc(5,6,add)
calc(78.0,55.0,sub)
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
func add(x,y int)int{
	fmt.Printf("%d + %d =%d\n",x,y,x+y)
	return x+y
}
func sub(x,y int)int{
	fmt.Printf("%d - %d =%d\n",x,y,x-y)
	return x-y
}
func calc(x,y int,op func(int,int)int)int{
	return op(x,y)
}

