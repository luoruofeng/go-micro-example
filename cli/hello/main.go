package main

import (

	"hello/handler"
	pb "hello/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "hello"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterHelloHandler(srv.Server(), new(handler.Hello)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
