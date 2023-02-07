package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/luoruofeng/go-micro-example/consul/myservice/config"
	"github.com/luoruofeng/go-micro-example/consul/myservice/handler"
	pb "github.com/luoruofeng/go-micro-example/consul/myservice/proto"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
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

	// init config
	if err := config.Init(consul_addr); err != nil {
		panic(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)

	//setup micro service
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(fmt.Sprintf(":%d", port)),
		micro.Registry(
			consul.NewRegistry(
				registry.Addrs(consul_addr),
			)),

		micro.RegisterTTL(20*time.Second),
		micro.RegisterInterval(10*time.Second),
	)

	// Register handler
	if err := pb.RegisterMyserviceHandler(srv.Server(), new(handler.Myservice)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	go func() {
		time.Sleep(time.Second * 3)
		// LoadConfig("./config.json")
		config.ConfigSet("abc", "123")
		fmt.Println(config.ConfigGet("abc"))

		config.Config.Get("mysql").Scan(&config.MysqlCnf)
		config.Config.Get("log").Scan(&config.LogCnf)
		fmt.Println(config.MysqlCnf)
		fmt.Println(config.LogCnf)
	}()

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
