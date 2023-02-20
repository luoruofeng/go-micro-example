package handler

import (
	"context"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "cservice/proto"
)

type Cservice struct {
	BService pb.BserviceService
}

func (e *Cservice) CMethod(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	logger.Infof("Cservice.CMethod()")
	e.BService.BMethod(context.TODO(), &emptypb.Empty{})
	return nil
}
