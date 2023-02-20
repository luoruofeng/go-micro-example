package handler

import (
	"context"

	"go-micro.dev/v4/logger"

	pb "aservice/proto"
)

type Aservice struct{}

func (e *Aservice) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	logger.Infof("Received Aservice.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}
