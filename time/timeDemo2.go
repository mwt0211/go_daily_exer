package main

import (
	"errors"
	"fmt"
	"time"
)

//解析字符串时间的类型
func main() {
	//拿到时间区域
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		errors.New("解析错误")
		return
	}
	//解析时间
	timeStr := "2020/12/05 22:22:22"
	parse, err := time.Parse("2006/01/02 15:04:05", timeStr)
	if err != nil {
		fmt.Printf(" time parse failed,err:%s\n", err)
	}
	fmt.Println(parse)
	parseInLocation, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, loc)
	if err != nil {
		fmt.Printf(" time ParseInLocation failed,err:%s\n", err)
	}
	fmt.Println(parseInLocation)

}
