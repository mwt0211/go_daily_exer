package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"mwt.com/go_daily_exer/GRPC/proto/user"
	"net"
	"net/http"
	"sync"
)

//1.取server
type UserService struct {
	user.UnimplementedUserServiceServer
}

//2.挂方法
func (U *UserService) SearchUser(ctx context.Context, req *user.UserRequest) (rep *user.UserResponse, err error) {
	name := req.GetName()
	uresp := &user.UserResponse{Name: "我收到了" + name + "的消息", Age: req.GetAge(), Addresses: req.GetAddresses()}
	return uresp, nil
}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go registerGateWay(&wg)
	go registerGRPC(&wg)
	wg.Wait()

}

func registerGRPC(sg *sync.WaitGroup) {
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal("注册服务前准备工作失败", err)
	}
	server := grpc.NewServer()
	user.RegisterUserServiceServer(server, &UserService{})
	//4.建立监听
	server.Serve(listen)
	fmt.Println("建立监听成功")
	sg.Done()
}

func registerGateWay(sg *sync.WaitGroup) {
	//创建客户端
	conn, _ := grpc.DialContext(
		context.Background(),
		"127.0.0.1:8081",
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	//2.创建mux
	mux := runtime.NewServeMux() //对外开放的mux
	//3.创建http服务器
	gwServer := &http.Server{
		Handler: mux,
		Addr:    ":8099",
	}
	//4.注册
	err2 := user.RegisterUserServiceHandler(context.Background(), mux, conn)
	if err2 != nil {
		log.Fatal("RegisterUserServiceHandler方法   注册失败")
	}
	gwServer.ListenAndServe()
	sg.Done()
}
