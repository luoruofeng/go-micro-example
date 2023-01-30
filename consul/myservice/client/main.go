package main

import (
	"fmt"

	"context"

	"github.com/go-micro/plugins/v4/registry/consul"

	pb "github.com/luoruofeng/go-micro-example/consul/myservice/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/registry"
)

func main() {
	service := micro.NewService(
		micro.Client(client.DefaultClient),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs("127.0.0.1:8500"),
			)),
	)
	service.Init()

	client := pb.NewMyserviceService("myservice", service.Client())
	rep, err := client.Call(context.Background(), &pb.CallRequest{
		Name: "luoruofeng",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rep.Msg)
}
