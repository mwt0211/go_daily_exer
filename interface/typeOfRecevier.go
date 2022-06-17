package main

import (
	"fmt"
)

type person1 []struct {
	name string
	age  int
}

//值接收和指针接收的区别
//接口的嵌套
type action interface {
	mover
	speaker
}
type person struct {
	name string
	age  int8
}
type mover interface {
	move()
}
type speaker interface {
	speak()
}

//func (p person) move() {//值传递
//	fmt.Printf("年龄%d,姓名为:%s的人在跑步  其类型为值传递\n", p.age, p.name)
//}
func (p *person) move() {
	fmt.Printf("年龄%d,姓名为:%s的人在跑步  其类型为指针 传递\n", p.age, p.name)
}
func (p *person) speak() {
	fmt.Printf("年龄%d,姓名为:%s的人在说话  其类型为指针 传递\n", p.age, p.name)
}
func main() {
	var m action
	//s := person{
	//	//值类型的person
	//	name: "张三",
	//	age:  18,
	//}

	p1 := &person{
		name: "赵四",
		age:  20,
	}
	m = p1
	fmt.Printf("类型为%T,%v\n", m, m)
	m.move()
	m.speak()
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
	var p2 = person1{{
		name: "陈晓瑞",
		age:  18,
	}, {
		name: "HPP",
		age:  20,
	},
	}
	x = p2
	for _, value := range p2 {
		fmt.Printf("type: %T,  value:%v\n", value, value)
	}
}
