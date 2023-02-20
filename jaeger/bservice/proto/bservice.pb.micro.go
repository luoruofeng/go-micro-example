// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/bservice.proto

package bservice

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Bservice service

func NewBserviceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Bservice service

type BserviceService interface {
	BMethod(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error)
}

type bserviceService struct {
	c    client.Client
	name string
}

func NewBserviceService(name string, c client.Client) BserviceService {
	return &bserviceService{
		c:    c,
		name: name,
	}
}

func (c *bserviceService) BMethod(ctx context.Context, in *emptypb.Empty, opts ...client.CallOption) (*emptypb.Empty, error) {
	req := c.c.NewRequest(c.name, "Bservice.BMethod", in)
	out := new(emptypb.Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Bservice service

type BserviceHandler interface {
	BMethod(context.Context, *emptypb.Empty, *emptypb.Empty) error
}

func RegisterBserviceHandler(s server.Server, hdlr BserviceHandler, opts ...server.HandlerOption) error {
	type bservice interface {
		BMethod(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error
	}
	type Bservice struct {
		bservice
	}
	h := &bserviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Bservice{h}, opts...))
}

type bserviceHandler struct {
	BserviceHandler
}

func (h *bserviceHandler) BMethod(ctx context.Context, in *emptypb.Empty, out *emptypb.Empty) error {
	return h.BserviceHandler.BMethod(ctx, in, out)
}

// Api Endpoints for Aservice service

func NewAserviceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Aservice service

type AserviceService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type aserviceService struct {
	c    client.Client
	name string
}

func NewAserviceService(name string, c client.Client) AserviceService {
	return &aserviceService{
		c:    c,
		name: name,
	}
}

func (c *aserviceService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Aservice.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Aservice service

type AserviceHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterAserviceHandler(s server.Server, hdlr AserviceHandler, opts ...server.HandlerOption) error {
	type aservice interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Aservice struct {
		aservice
	}
	h := &aserviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Aservice{h}, opts...))
}

type aserviceHandler struct {
	AserviceHandler
}

func (h *aserviceHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.AserviceHandler.Call(ctx, in, out)
}