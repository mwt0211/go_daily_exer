package main

import "fmt"

func main() {
	//数组初始化
	//1.定义时使用初始值列表
	var cityarray = [4]string{"北京", "上海", "河南", "浙江"}
	//	fmt.Println(cityarray)   //  [北京 上海 河南 浙江]
	//2.编译器推导数组的长度
	//var numberarray=[...]int{1,2,3,4,5}
	//	fmt.Println(numberarray)  //  [1 2 3 4 5]
	//	fmt.Printf("数组的长度为%d,分别为:%v\n",len(numberarray),numberarray) //数组的长度为5,分别为:[1 2 3 4 5]
	//3.使用索引值方式初始化
	var intArray = [...]int{1: 99, 5: 520, 2: 1314}
	//	fmt.Printf("数组的长度为%d,分别为:%v\n",len(intArray),intArray) //数组的长度为6,分别为:[0 99 1314 0 0 520]
	//数组遍历
	//1.使用for遍历
	for i := 0; i < len(intArray); i++ {
		fmt.Printf("第%d个元素为%d \n", i+1, intArray[i])
	}
	//2.使用range遍历
	for index, value := range cityarray {
		fmt.Printf("索引为:%d,对应的城市的值为:%s \n", index, value)
	}

}
