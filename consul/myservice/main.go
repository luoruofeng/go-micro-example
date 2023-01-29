package main

import (
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/luorufoeng/go-micro-example/consul/myservice/handler"
	pb "github.com/luorufoeng/go-micro-example/consul/myservice/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "myservice"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Registry(consul.NewRegistry()),
		micro.Address(":8888"),
	)

	// Register handler
	if err := pb.RegisterMyserviceHandler(srv.Server(), new(handler.Myservice)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
