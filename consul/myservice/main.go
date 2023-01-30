package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/luoruofeng/go-micro-example/consul/myservice/handler"
	pb "github.com/luoruofeng/go-micro-example/consul/myservice/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

var (
	service = "myservice"
	version = "latest"
)

func main() {
	myservice_port := os.Getenv("myservice_port")
	consul_addr := os.Getenv("consul_addr")

	//listen random port
	port := 0
	if myservice_port != "" {
		port, _ = strconv.Atoi(myservice_port)
	}

	if consul_addr == "" {
		consul_addr = "127.0.0.1:8500"
	}

	// Create service
	srv := micro.NewService()

	//setup micro service
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(fmt.Sprintf(":%d", port)),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs(consul_addr),
			)),

		micro.RegisterTTL(time.Minute*30),
		micro.RegisterInterval(time.Minute*10),
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
