// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/license/key.proto

package license

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/asim/go-micro/v3/api"
	client "github.com/asim/go-micro/v3/client"
	server "github.com/asim/go-micro/v3/server"
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

// Api Endpoints for Key service

func NewKeyEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Key service

type KeyService interface {
	// 生成
	Generate(ctx context.Context, in *KeyGenerateRequest, opts ...client.CallOption) (*KeyGenerateResponse, error)
	// 激活
	Activate(ctx context.Context, in *KeyActivateRequest, opts ...client.CallOption) (*KeyActivateResponse, error)
	// 更新
	Update(ctx context.Context, in *KeyUpdateRequest, opts ...client.CallOption) (*UuidResponse, error)
	// 获取
	Get(ctx context.Context, in *KeyGetRequest, opts ...client.CallOption) (*KeyGetResponse, error)
	// 列举
	List(ctx context.Context, in *KeyListRequest, opts ...client.CallOption) (*KeyListResponse, error)
	// 查询
	Search(ctx context.Context, in *KeySearchRequest, opts ...client.CallOption) (*KeyListResponse, error)
}

type keyService struct {
	c    client.Client
	name string
}

func NewKeyService(name string, c client.Client) KeyService {
	return &keyService{
		c:    c,
		name: name,
	}
}

func (c *keyService) Generate(ctx context.Context, in *KeyGenerateRequest, opts ...client.CallOption) (*KeyGenerateResponse, error) {
	req := c.c.NewRequest(c.name, "Key.Generate", in)
	out := new(KeyGenerateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyService) Activate(ctx context.Context, in *KeyActivateRequest, opts ...client.CallOption) (*KeyActivateResponse, error) {
	req := c.c.NewRequest(c.name, "Key.Activate", in)
	out := new(KeyActivateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyService) Update(ctx context.Context, in *KeyUpdateRequest, opts ...client.CallOption) (*UuidResponse, error) {
	req := c.c.NewRequest(c.name, "Key.Update", in)
	out := new(UuidResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyService) Get(ctx context.Context, in *KeyGetRequest, opts ...client.CallOption) (*KeyGetResponse, error) {
	req := c.c.NewRequest(c.name, "Key.Get", in)
	out := new(KeyGetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyService) List(ctx context.Context, in *KeyListRequest, opts ...client.CallOption) (*KeyListResponse, error) {
	req := c.c.NewRequest(c.name, "Key.List", in)
	out := new(KeyListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyService) Search(ctx context.Context, in *KeySearchRequest, opts ...client.CallOption) (*KeyListResponse, error) {
	req := c.c.NewRequest(c.name, "Key.Search", in)
	out := new(KeyListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Key service

type KeyHandler interface {
	// 生成
	Generate(context.Context, *KeyGenerateRequest, *KeyGenerateResponse) error
	// 激活
	Activate(context.Context, *KeyActivateRequest, *KeyActivateResponse) error
	// 更新
	Update(context.Context, *KeyUpdateRequest, *UuidResponse) error
	// 获取
	Get(context.Context, *KeyGetRequest, *KeyGetResponse) error
	// 列举
	List(context.Context, *KeyListRequest, *KeyListResponse) error
	// 查询
	Search(context.Context, *KeySearchRequest, *KeyListResponse) error
}

func RegisterKeyHandler(s server.Server, hdlr KeyHandler, opts ...server.HandlerOption) error {
	type key interface {
		Generate(ctx context.Context, in *KeyGenerateRequest, out *KeyGenerateResponse) error
		Activate(ctx context.Context, in *KeyActivateRequest, out *KeyActivateResponse) error
		Update(ctx context.Context, in *KeyUpdateRequest, out *UuidResponse) error
		Get(ctx context.Context, in *KeyGetRequest, out *KeyGetResponse) error
		List(ctx context.Context, in *KeyListRequest, out *KeyListResponse) error
		Search(ctx context.Context, in *KeySearchRequest, out *KeyListResponse) error
	}
	type Key struct {
		key
	}
	h := &keyHandler{hdlr}
	return s.Handle(s.NewHandler(&Key{h}, opts...))
}

type keyHandler struct {
	KeyHandler
}

func (h *keyHandler) Generate(ctx context.Context, in *KeyGenerateRequest, out *KeyGenerateResponse) error {
	return h.KeyHandler.Generate(ctx, in, out)
}

func (h *keyHandler) Activate(ctx context.Context, in *KeyActivateRequest, out *KeyActivateResponse) error {
	return h.KeyHandler.Activate(ctx, in, out)
}

func (h *keyHandler) Update(ctx context.Context, in *KeyUpdateRequest, out *UuidResponse) error {
	return h.KeyHandler.Update(ctx, in, out)
}

func (h *keyHandler) Get(ctx context.Context, in *KeyGetRequest, out *KeyGetResponse) error {
	return h.KeyHandler.Get(ctx, in, out)
}

func (h *keyHandler) List(ctx context.Context, in *KeyListRequest, out *KeyListResponse) error {
	return h.KeyHandler.List(ctx, in, out)
}

func (h *keyHandler) Search(ctx context.Context, in *KeySearchRequest, out *KeyListResponse) error {
	return h.KeyHandler.Search(ctx, in, out)
}
