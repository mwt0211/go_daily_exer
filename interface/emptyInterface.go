package main

import (
	"fmt"
)

type person1 []struct {
	name string
	age  int
}

//空接口
func main() {
	//定义一个空接口
	var x interface{}
	str := "hello 瑞瑞"
	x = str
	fmt.Printf("type: %T,  value:%v\n", x, x)
	num := 100
	x = num
	fmt.Printf("type: %T,  value:%v\n", x, x)
	bol := true
	x = bol
	fmt.Printf("type: %T,  value:%v\n", x, x)
	var p1 = person1{{
		name: "陈晓瑞",
		age:  18,
	}, {
		name: "HPP",
		age:  20,
	},
	}
	x = p1
	for _, value := range p1 {
		fmt.Printf("type: %T,  value:%v\n", value, value)
	}

}
