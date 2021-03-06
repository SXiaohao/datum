// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: sms.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for SmsService service

func NewSmsServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for SmsService service

type SmsService interface {
	SendRegisterCode(ctx context.Context, in *SendRegisterCodeRequest, opts ...client.CallOption) (*SendRegisterCodeResponse, error)
	SendRestPwdCode(ctx context.Context, in *SendRestPwdCodeRequest, opts ...client.CallOption) (*SendRestPwdCodeResponse, error)
}

type smsService struct {
	c    client.Client
	name string
}

func NewSmsService(name string, c client.Client) SmsService {
	return &smsService{
		c:    c,
		name: name,
	}
}

func (c *smsService) SendRegisterCode(ctx context.Context, in *SendRegisterCodeRequest, opts ...client.CallOption) (*SendRegisterCodeResponse, error) {
	req := c.c.NewRequest(c.name, "SmsService.SendRegisterCode", in)
	out := new(SendRegisterCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *smsService) SendRestPwdCode(ctx context.Context, in *SendRestPwdCodeRequest, opts ...client.CallOption) (*SendRestPwdCodeResponse, error) {
	req := c.c.NewRequest(c.name, "SmsService.SendRestPwdCode", in)
	out := new(SendRestPwdCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SmsService service

type SmsServiceHandler interface {
	SendRegisterCode(context.Context, *SendRegisterCodeRequest, *SendRegisterCodeResponse) error
	SendRestPwdCode(context.Context, *SendRestPwdCodeRequest, *SendRestPwdCodeResponse) error
}

func RegisterSmsServiceHandler(s server.Server, hdlr SmsServiceHandler, opts ...server.HandlerOption) error {
	type smsService interface {
		SendRegisterCode(ctx context.Context, in *SendRegisterCodeRequest, out *SendRegisterCodeResponse) error
		SendRestPwdCode(ctx context.Context, in *SendRestPwdCodeRequest, out *SendRestPwdCodeResponse) error
	}
	type SmsService struct {
		smsService
	}
	h := &smsServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&SmsService{h}, opts...))
}

type smsServiceHandler struct {
	SmsServiceHandler
}

func (h *smsServiceHandler) SendRegisterCode(ctx context.Context, in *SendRegisterCodeRequest, out *SendRegisterCodeResponse) error {
	return h.SmsServiceHandler.SendRegisterCode(ctx, in, out)
}

func (h *smsServiceHandler) SendRestPwdCode(ctx context.Context, in *SendRestPwdCodeRequest, out *SendRestPwdCodeResponse) error {
	return h.SmsServiceHandler.SendRestPwdCode(ctx, in, out)
}
