// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/myservice.proto

package myservice

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
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

// Api Endpoints for Myservice service

func NewMyserviceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Myservice service

type MyserviceService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
}

type myserviceService struct {
	c    client.Client
	name string
}

func NewMyserviceService(name string, c client.Client) MyserviceService {
	return &myserviceService{
		c:    c,
		name: name,
	}
}

func (c *myserviceService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "Myservice.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Myservice service

type MyserviceHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
}

func RegisterMyserviceHandler(s server.Server, hdlr MyserviceHandler, opts ...server.HandlerOption) error {
	type myservice interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
	}
	type Myservice struct {
		myservice
	}
	h := &myserviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Myservice{h}, opts...))
}

type myserviceHandler struct {
	MyserviceHandler
}

func (h *myserviceHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.MyserviceHandler.Call(ctx, in, out)
}