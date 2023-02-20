// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/health.proto

package aservice

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

// Api Endpoints for Health service

func NewHealthEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Health service

type HealthService interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...client.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...client.CallOption) (Health_WatchService, error)
}

type healthService struct {
	c    client.Client
	name string
}

func NewHealthService(name string, c client.Client) HealthService {
	return &healthService{
		c:    c,
		name: name,
	}
}

func (c *healthService) Check(ctx context.Context, in *HealthCheckRequest, opts ...client.CallOption) (*HealthCheckResponse, error) {
	req := c.c.NewRequest(c.name, "Health.Check", in)
	out := new(HealthCheckResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthService) Watch(ctx context.Context, in *HealthCheckRequest, opts ...client.CallOption) (Health_WatchService, error) {
	req := c.c.NewRequest(c.name, "Health.Watch", &HealthCheckRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &healthServiceWatch{stream}, nil
}

type Health_WatchService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*HealthCheckResponse, error)
}

type healthServiceWatch struct {
	stream client.Stream
}

func (x *healthServiceWatch) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *healthServiceWatch) Close() error {
	return x.stream.Close()
}

func (x *healthServiceWatch) Context() context.Context {
	return x.stream.Context()
}

func (x *healthServiceWatch) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *healthServiceWatch) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *healthServiceWatch) Recv() (*HealthCheckResponse, error) {
	m := new(HealthCheckResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Health service

type HealthHandler interface {
	Check(context.Context, *HealthCheckRequest, *HealthCheckResponse) error
	Watch(context.Context, *HealthCheckRequest, Health_WatchStream) error
}

func RegisterHealthHandler(s server.Server, hdlr HealthHandler, opts ...server.HandlerOption) error {
	type health interface {
		Check(ctx context.Context, in *HealthCheckRequest, out *HealthCheckResponse) error
		Watch(ctx context.Context, stream server.Stream) error
	}
	type Health struct {
		health
	}
	h := &healthHandler{hdlr}
	return s.Handle(s.NewHandler(&Health{h}, opts...))
}

type healthHandler struct {
	HealthHandler
}

func (h *healthHandler) Check(ctx context.Context, in *HealthCheckRequest, out *HealthCheckResponse) error {
	return h.HealthHandler.Check(ctx, in, out)
}

func (h *healthHandler) Watch(ctx context.Context, stream server.Stream) error {
	m := new(HealthCheckRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.HealthHandler.Watch(ctx, m, &healthWatchStream{stream})
}

type Health_WatchStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*HealthCheckResponse) error
}

type healthWatchStream struct {
	stream server.Stream
}

func (x *healthWatchStream) Close() error {
	return x.stream.Close()
}

func (x *healthWatchStream) Context() context.Context {
	return x.stream.Context()
}

func (x *healthWatchStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *healthWatchStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *healthWatchStream) Send(m *HealthCheckResponse) error {
	return x.stream.Send(m)
}
