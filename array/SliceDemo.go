package main

import "fmt"

//Slice切片
func main() {
	/*	//1.基于数组定义切片
		var a=[5]string{"重庆","北京","上海","河南","浙江"}
		b:=a[1:4] //包头不包尾
		fmt.Printf("切片b类型为:%T 其值为:%v\n",b,b)
		//2.切片再次切片
		c:=b[:]
		fmt.Printf("切片c类型为:%T 其值为:%v\n",c,c)
		//3.通过make函数构造切片
		d:=make([]int,5,100)
		fmt.Printf("切片d类型为%T,长度为:%d,其值为:%v\n",d,len(d),d)
		//通过cap()获取切片的容量(最多能存放的个数)
		fmt.Printf("长度为:%d,容量为:%d\n",len(d),cap(d))*/
	//切片的赋值拷贝
	intArr := make([]int, 3)
	intArr_b := intArr
	intArr_b[1] = 222
	fmt.Println(intArr)

}
