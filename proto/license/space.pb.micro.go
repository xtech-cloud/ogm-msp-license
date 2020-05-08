// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/license/space.proto

package license

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Space service

type SpaceService interface {
	// 创建
	Create(ctx context.Context, in *SpaceCreateRequest, opts ...client.CallOption) (*BlankResponse, error)
	// 查询
	Query(ctx context.Context, in *SpaceQueryRequest, opts ...client.CallOption) (*SpaceQueryResponse, error)
	// 列举
	List(ctx context.Context, in *SpaceListRequest, opts ...client.CallOption) (*SpaceListResponse, error)
}

type spaceService struct {
	c    client.Client
	name string
}

func NewSpaceService(name string, c client.Client) SpaceService {
	return &spaceService{
		c:    c,
		name: name,
	}
}

func (c *spaceService) Create(ctx context.Context, in *SpaceCreateRequest, opts ...client.CallOption) (*BlankResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Create", in)
	out := new(BlankResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) Query(ctx context.Context, in *SpaceQueryRequest, opts ...client.CallOption) (*SpaceQueryResponse, error) {
	req := c.c.NewRequest(c.name, "Space.Query", in)
	out := new(SpaceQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spaceService) List(ctx context.Context, in *SpaceListRequest, opts ...client.CallOption) (*SpaceListResponse, error) {
	req := c.c.NewRequest(c.name, "Space.List", in)
	out := new(SpaceListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Space service

type SpaceHandler interface {
	// 创建
	Create(context.Context, *SpaceCreateRequest, *BlankResponse) error
	// 查询
	Query(context.Context, *SpaceQueryRequest, *SpaceQueryResponse) error
	// 列举
	List(context.Context, *SpaceListRequest, *SpaceListResponse) error
}

func RegisterSpaceHandler(s server.Server, hdlr SpaceHandler, opts ...server.HandlerOption) error {
	type space interface {
		Create(ctx context.Context, in *SpaceCreateRequest, out *BlankResponse) error
		Query(ctx context.Context, in *SpaceQueryRequest, out *SpaceQueryResponse) error
		List(ctx context.Context, in *SpaceListRequest, out *SpaceListResponse) error
	}
	type Space struct {
		space
	}
	h := &spaceHandler{hdlr}
	return s.Handle(s.NewHandler(&Space{h}, opts...))
}

type spaceHandler struct {
	SpaceHandler
}

func (h *spaceHandler) Create(ctx context.Context, in *SpaceCreateRequest, out *BlankResponse) error {
	return h.SpaceHandler.Create(ctx, in, out)
}

func (h *spaceHandler) Query(ctx context.Context, in *SpaceQueryRequest, out *SpaceQueryResponse) error {
	return h.SpaceHandler.Query(ctx, in, out)
}

func (h *spaceHandler) List(ctx context.Context, in *SpaceListRequest, out *SpaceListResponse) error {
	return h.SpaceHandler.List(ctx, in, out)
}
