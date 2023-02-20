package main

import (
	"aservice/handler"
	pb "aservice/proto"

	ot "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/cli/debug/trace/jaeger"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
)

var (
	service = "aservice"
	version = "latest"
)

func main() {
	// Create tracer
	tracer, closer, err := jaeger.NewTracer(
		jaeger.Name(service),
		jaeger.FromEnv(true),
		jaeger.GlobalTracer(true),
	)
	if err != nil {
		logger.Fatal(err)
	}
	defer closer.Close()

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		// micro.WrapCall(ot.NewCallWrapper(tracer)),
		// micro.WrapClient(ot.NewClientWrapper(tracer)),
		// micro.WrapSubscriber(ot.NewSubscriberWrapper(tracer)),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterAserviceHandler(srv.Server(), new(handler.Aservice)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
