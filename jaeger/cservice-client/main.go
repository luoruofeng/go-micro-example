package main

import (
	"context"
	pb "cservice-client/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/go-micro/plugins/v4/client/grpc"
)

var (
	service = "cservice"
	version = "latest"
)

func main() {
	// Create service

	srv := micro.NewService(
		micro.Client(grpc.NewClient()),
	)

	srv.Init()

	// Create client
	c := pb.NewCserviceService(service, srv.Client())

	// Call service
	_, err := c.CMethod(context.Background(), &emptypb.Empty{})
	if err != nil {
		logger.Fatal(err)
	}

}
