package main

import "fmt"

func main() {
	var a map[int]string
	fmt.Println(a == nil)
	//map初始化
	a = make(map[int]string, 4)
	fmt.Println(a == nil)
	//map中添加键值对
	a[1] = "沙特"
	a[2] = "沙和尚"
	a[3] = "孙悟空"
	a[4] = "唐僧"
	a[5] = "唐僧1"
	a[6] = "唐僧2"
	fmt.Printf("a:%#v\n", a)
	//声明map的同时完成初始化
	b := map[string]string{"001": "张三", "002": "李四", "003": "小红", "004": "小明", "005": "小李"}
	fmt.Printf("b:%#v\n", b)
	//判断某个键是否存在
	score := make(map[string]int, 4)
	score["小明"] = 80
	score["小王"] = 81
	score["小李"] = 70
	score["小宋"] = 88
	score["小马"] = 100
	score["小陈"] = 60
	score["小陈"] = 70
	value, ok := score["小宋"]
	if ok {
		fmt.Printf("小宋的分数为:%d", value)
	} else {
		fmt.Printf("查无此人")
	}
	//遍历key与value
	/*	for key, value := range score {
		fmt.Printf("学生的名字是:%v, 分数为:%d\n",key,value)
	}*/
	//删除某个键值对
	fmt.Printf("%v,长度为%d\n", score, len(score))
	delete(score, "小陈")
	fmt.Printf("%v,长度为%d\n", score, len(score))

}
