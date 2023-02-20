package handler

import (
	"context"

	"go-micro.dev/v4/logger"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "bservice/proto"
)

type Bservice struct {
	Aservice pb.AserviceService
}

func (e *Bservice) BMethod(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	logger.Infof("BService.BMethod()")
	rep, _ := e.Aservice.Call(context.TODO(), &pb.CallRequest{Name: "luo"})
	logger.Infof("call Aservice.Call() get response:%v\n", rep.Msg)
	return nil
}
