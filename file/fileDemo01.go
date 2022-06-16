package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

//文件操作
func main() {
	//of, err := os.Open("./exer.txt")//相对路径
	//打开文件

	//ReadJsonFile()
	ReadCycle()
	//ReadTxt()

}

//读取json文件
func ReadJsonFile() error {
	jsfile, err := os.Open("C:\\Users\\ada\\Desktop\\exer_about_Go\\go_daily_exer\\file\\as.json")

	defer jsfile.Close()
	var loggerConfig map[string]string
	byteValue, _ := ioutil.ReadAll(jsfile)
	json.Unmarshal([]byte(byteValue), &loggerConfig)
	for key, loggerConfig_value := range loggerConfig {
		fmt.Printf(key + ":" + loggerConfig_value + "\n")
	}
	return err
}

//读取text文件
func ReadTxt() error {
	//读取exer.txt文件
	file, err := os.Open("C:\\Users\\ada\\Desktop\\exer_about_Go\\go_daily_exer\\file\\exer.txt") //绝对路径
	if err != nil {
		fmt.Printf("file open failed,err:%v", err)
	}
	defer file.Close() //关闭文件
	readAll, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("文件读取失败,err:%s", err)
	}
	fmt.Println(string(readAll))
	return err
}

//循环读取文件
func ReadCycle() error {
	file, err := os.Open("file/xunhuan.txt")
	if err != nil {
		fmt.Printf("file read failed,err:%s\n", err)
	}
	defer file.Close()
	var outfile []string
	for {
		var tmp = make([]byte, 128)
		_, err := file.Read(tmp)
		if err == io.EOF {

			break
		}
		if err != nil {
			fmt.Printf("file read failed,err:%s\n", err)
		}
		//fmt.Println(string(tmp))
		outfile = append(outfile, string(tmp))
	}
	fmt.Println(outfile)
	for _, s := range outfile {
		fmt.Printf(s)
	}

	return err
}
