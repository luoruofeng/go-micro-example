package main

import (
	"github.com/luorufoeng/go-micro-example/cli/helloworld/handler"
	pb "github.com/luorufoeng/go-micro-example/cli/helloworld/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "helloworld"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService()
	srv.Init(
		micro.Name(service),
		micro.Version(version),
		micro.Address(":8081"),
	)

	// Register handler
	if err := pb.RegisterHelloworldHandler(srv.Server(), new(handler.Helloworld)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
