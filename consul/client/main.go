package main

import (
	"fmt"
	"os"

	"context"

	"github.com/go-micro/plugins/v4/registry/consul"

	pb "github.com/luoruofeng/go-micro-example/consul/myservice/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
)

var (
	ServiceName = "myservice"
)

func main() {
	fmt.Println("---myclient start---")
	consul_addr_arg := os.Getenv("consul_addr")
	//listen random port
	consul_addr := "127.0.0.1:8500"
	if consul_addr_arg != "" {
		consul_addr = consul_addr_arg
	}
	fmt.Println(consul_addr)

	service := micro.NewService(
		micro.Client(grpcc.NewClient()),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs(consul_addr),
			)),
	)
	service.Init()
	client := service.Client()

	myService := pb.NewMyserviceService(ServiceName, client)

	callReq := &pb.CallRequest{Name: "luoruofeng"}
	rep, err := myService.Call(context.Background(), callReq)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rep.Msg)
}
