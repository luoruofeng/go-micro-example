package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "github.com/luoruofeng/go-micro-example/consul/myservice/proto"
)

type Myservice struct{}

func (e *Myservice) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Myservice.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
