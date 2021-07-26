// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/license/certificate.proto

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

// Api Endpoints for Certificate service

func NewCertificateEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Certificate service

type CertificateService interface {
	// 获取
	Fetch(ctx context.Context, in *CerFetchRequest, opts ...client.CallOption) (*CerFetchResponse, error)
	// 拉取
	Pull(ctx context.Context, in *CerPullRequest, opts ...client.CallOption) (*CerPullResponse, error)
	// 列举
	List(ctx context.Context, in *CerListRequest, opts ...client.CallOption) (*CerListResponse, error)
}

type certificateService struct {
	c    client.Client
	name string
}

func NewCertificateService(name string, c client.Client) CertificateService {
	return &certificateService{
		c:    c,
		name: name,
	}
}

func (c *certificateService) Fetch(ctx context.Context, in *CerFetchRequest, opts ...client.CallOption) (*CerFetchResponse, error) {
	req := c.c.NewRequest(c.name, "Certificate.Fetch", in)
	out := new(CerFetchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateService) Pull(ctx context.Context, in *CerPullRequest, opts ...client.CallOption) (*CerPullResponse, error) {
	req := c.c.NewRequest(c.name, "Certificate.Pull", in)
	out := new(CerPullResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateService) List(ctx context.Context, in *CerListRequest, opts ...client.CallOption) (*CerListResponse, error) {
	req := c.c.NewRequest(c.name, "Certificate.List", in)
	out := new(CerListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Certificate service

type CertificateHandler interface {
	// 获取
	Fetch(context.Context, *CerFetchRequest, *CerFetchResponse) error
	// 拉取
	Pull(context.Context, *CerPullRequest, *CerPullResponse) error
	// 列举
	List(context.Context, *CerListRequest, *CerListResponse) error
}

func RegisterCertificateHandler(s server.Server, hdlr CertificateHandler, opts ...server.HandlerOption) error {
	type certificate interface {
		Fetch(ctx context.Context, in *CerFetchRequest, out *CerFetchResponse) error
		Pull(ctx context.Context, in *CerPullRequest, out *CerPullResponse) error
		List(ctx context.Context, in *CerListRequest, out *CerListResponse) error
	}
	type Certificate struct {
		certificate
	}
	h := &certificateHandler{hdlr}
	return s.Handle(s.NewHandler(&Certificate{h}, opts...))
}

type certificateHandler struct {
	CertificateHandler
}

func (h *certificateHandler) Fetch(ctx context.Context, in *CerFetchRequest, out *CerFetchResponse) error {
	return h.CertificateHandler.Fetch(ctx, in, out)
}

func (h *certificateHandler) Pull(ctx context.Context, in *CerPullRequest, out *CerPullResponse) error {
	return h.CertificateHandler.Pull(ctx, in, out)
}

func (h *certificateHandler) List(ctx context.Context, in *CerListRequest, out *CerListResponse) error {
	return h.CertificateHandler.List(ctx, in, out)
}
