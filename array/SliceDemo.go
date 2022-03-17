package main

import "fmt"

//Slice切片为引用类型
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
	/*	intArr := make([]int, 3)
		intArr_b := intArr
		intArr_b[1] = 222
		fmt.Println(intArr)*/
	//Slice 的扩容
	//note:
	var numberSlice []int //此时并未申请内存
	for i := 0; i < 16; i++ {
		numberSlice = append(numberSlice, i)
		fmt.Printf("%v,len:%d,cap:%d,ptr:%p\n", numberSlice, len(numberSlice), cap(numberSlice), numberSlice)
	}
	//一次追加多个元素
	var citySlice = []string{"北京", "上海", "南京", "广东", "河南"}
	var citySlice1 = []string{"新乡", "周口", "洛阳", "郑州", "南阳"}
	citySlice = append(citySlice, citySlice1...)
	//for _, value := range citySlice {
	//	fmt.Println(value)
	//}
	fmt.Println(citySlice)

}
