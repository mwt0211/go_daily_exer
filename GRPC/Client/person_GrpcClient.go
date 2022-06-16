package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"mwt.com/go_daily_exer/GRPC/proto/person"
	"sync"
	"time"
)

func main() {
	//1.创建连接
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal("客户端创建连接失败", err)
	}
	//2.创建客户端
	personServiceClient := person.NewPersonServiceClient(conn)
	//3.1调普通方法
	/*	searchPerson, err := personServiceClient.SearchPerson(context.Background(), &person.PersonRequest{Age: 18, Name: "张三"})
		if err != nil {
			log.Fatal("客户端调用方法SearchPerson失败")
		}
		fmt.Println(searchPerson)*/
	/*//3.2调用请求为stream的方法
	reqstream, err := personServiceClient.SearchPersonReqStream(context.Background())
	if err != nil {
		log.Fatal("客户端发送数据失败",err)
	}
	count:=0//便于记录客户端发送消息的次数
	for{
		if count>10{
			recv, _ := reqstream.CloseAndRecv()
			fmt.Println(recv)
			break
		}
		reqstream.Send(&person.PersonRequest{Name: "我是客户端发送进来的消息"})
		time.Sleep(time.Second)//每隔一秒发送一次消息
		count++
	}*/
	/*//3.3调用服务端为stream的方法
	resStream, err := personServiceClient.SearchPersonResStream(context.Background(), &person.PersonRequest{Name: "MWT"})
	if err != nil {
		log.Fatal("调用SearchPersonResStream方法异常")
	}
	for{
		recv, err := resStream.Recv()
		if err != nil {
			if err==io.EOF{
				log.Fatal("服务端发送数据完成:--->",err)
				break
			}
			log.Fatal("接收服务端发送的数据异常:--->",err)
			break
		}
		fmt.Println(recv)
	}*/
	//3.4客户端服务端均为stream
	bothStream, err := personServiceClient.SearchPersonBothStream(context.Background())
	if err != nil {
		log.Fatal("调用SearchPersonBothStream方法异常")
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	//客户端不断向服务端发送数据
	count := 0
	go func() {
		for {
			err2 := bothStream.Send(&person.PersonRequest{Name: "客户端向服务端不断发送数据"})
			if err2 != nil {
				if err2 == io.EOF {
					log.Fatal("客户端发送数据完成", err2)
					wg.Done()
					break
				}
				log.Fatal("客户端发送数据异常", err2)
				wg.Done()
				break
			}
			time.Sleep(time.Second)
			count++
			if count > 10 {
				break
			}
		}

	}()
	//客户端不断接收服务端发送的数据
	go func() {
		for {
			recv, err2 := bothStream.Recv()
			if err2 != nil {
				if err2 == io.EOF {
					log.Fatal("客户端接收数据完成", err2)
					wg.Done()
					break
				}
				log.Fatal("客户端接收数据异常", err2)
				wg.Done()
				break
			}
			fmt.Println(recv)
		}

	}()

	wg.Wait()
	//4.
}
