package main

import (
	"fmt"

	"context"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/luorufoeng/go-micro-example/consul/myservice/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
)

func main() {
	service := micro.NewService(
		micro.Client(client.DefaultClient()),
		micro.Registry(consul.NewRegistry()),
	)
	service.Init()

	client := proto.NewMyserviceService("", service.Client())
	rep, err := client.Call(context.Background(), &proto.CallRequest{
		Name: "luoruofeng",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rep.Msg)

}
