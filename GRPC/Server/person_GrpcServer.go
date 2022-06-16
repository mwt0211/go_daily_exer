package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mwt.com/go_daily_exer/GRPC/proto/person"
	"net"
	"time"
)

//1.取server
type PersonService struct {
	person.UnimplementedPersonServiceServer
}

//2.挂载方法
func (PersonSer *PersonService) SearchPerson(ctx context.Context, req *person.PersonRequest) (res *person.PersonResponse, err error) {
	//即刻响应
	name := req.GetName()
	age := req.GetAge()
	Response := &person.PersonResponse{Name: "我收到了" + name + "发送来的信息", Age: age}
	return Response, nil
}
func (PersonSer *PersonService) SearchPersonReqStream(req person.PersonService_SearchPersonReqStreamServer) error {
	//请求为stream

	for {
		recv, err := req.Recv() //不断读客户端发送的数据
		fmt.Println("客户端发送的数据为:", recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端传输数据完成")
				req.SendAndClose(&person.PersonResponse{Name: "数据传输完成"})
				break
			}
			log.Fatal("客户端发送stream数据,服务端准备接收数据失败", err)
		}
	}
	return nil
}
func (PersonSer *PersonService) SearchPersonResStream(req *person.PersonRequest, res person.PersonService_SearchPersonResStreamServer) error {
	//响应为stream
	//服务端不断发送请求
	name := req.Name
	count := 0
	for {
		if count > 10 {
			break
		}
		time.Sleep(time.Second)
		err := res.Send(&person.PersonResponse{Name: "我是服务端不断发送的请求" + name})
		count++
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务端发送数据完成")
			}
			fmt.Println("服务端发送数据异常终止")
		}
	}
	return nil
}
func (PersonSer *PersonService) SearchPersonBothStream(both person.PersonService_SearchPersonBothStreamServer) error {
	//请求响应均为stream,可用于即时聊天
	//服务端先接收信息,然后在发送信息
	str := make(chan string)
	go func() {
		for {
			recv, err := both.Recv()
			if err != nil {
				if err == io.EOF {
					str <- "服务端接收客户端发送的数据,接收完毕"
					fmt.Println("服务端接收客户端发送的数据,接收完毕", err)
					break
				}
				str <- "结束"
				fmt.Println("服务端接收客户端发送的数据,中途发生异常", err)
				break
			}
			str <- recv.Name
			fmt.Println("服务端接收客户端发送的数据为:", recv)
		}
	}()
	//服务端向客户端发送数据
	count := 0
	for {
		err := both.Send(&person.PersonResponse{Name: <-str})
		if err != nil {
			if err == io.EOF {
				fmt.Println("服务端向客户端发送数据完成", err)
				break
			}
			fmt.Println("服务端向客户端发送数据异常", err)
			break
		}
		count++
		time.Sleep(time.Second)
		if count > 10 {
			break
		}
	}
	return nil
}
func main() {
	//3.注册服务
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("注册服务前准备工作失败", err)
	}
	server := grpc.NewServer()
	person.RegisterPersonServiceServer(server, &PersonService{})
	//4.建立监听
	server.Serve(listen)
	fmt.Println("建立监听成功")
}
