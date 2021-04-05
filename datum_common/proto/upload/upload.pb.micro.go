// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: upload.proto

package upload

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

// Api Endpoints for Upload service

func NewUploadEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Upload service

type UploadService interface {
	GetPicturePath(ctx context.Context, in *UploadPictureRequest, opts ...client.CallOption) (*UploadPictureResponse, error)
	GetFilePath(ctx context.Context, in *UploadFileRequest, opts ...client.CallOption) (*UploadFileResponse, error)
	CheckText(ctx context.Context, in *CheckTextRequest, opts ...client.CallOption) (*CheckTextResponse, error)
}

type uploadService struct {
	c    client.Client
	name string
}

func NewUploadService(name string, c client.Client) UploadService {
	return &uploadService{
		c:    c,
		name: name,
	}
}

func (c *uploadService) GetPicturePath(ctx context.Context, in *UploadPictureRequest, opts ...client.CallOption) (*UploadPictureResponse, error) {
	req := c.c.NewRequest(c.name, "Upload.GetPicturePath", in)
	out := new(UploadPictureResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadService) GetFilePath(ctx context.Context, in *UploadFileRequest, opts ...client.CallOption) (*UploadFileResponse, error) {
	req := c.c.NewRequest(c.name, "Upload.GetFilePath", in)
	out := new(UploadFileResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *uploadService) CheckText(ctx context.Context, in *CheckTextRequest, opts ...client.CallOption) (*CheckTextResponse, error) {
	req := c.c.NewRequest(c.name, "Upload.CheckText", in)
	out := new(CheckTextResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Upload service

type UploadHandler interface {
	GetPicturePath(context.Context, *UploadPictureRequest, *UploadPictureResponse) error
	GetFilePath(context.Context, *UploadFileRequest, *UploadFileResponse) error
	CheckText(context.Context, *CheckTextRequest, *CheckTextResponse) error
}

func RegisterUploadHandler(s server.Server, hdlr UploadHandler, opts ...server.HandlerOption) error {
	type upload interface {
		GetPicturePath(ctx context.Context, in *UploadPictureRequest, out *UploadPictureResponse) error
		GetFilePath(ctx context.Context, in *UploadFileRequest, out *UploadFileResponse) error
		CheckText(ctx context.Context, in *CheckTextRequest, out *CheckTextResponse) error
	}
	type Upload struct {
		upload
	}
	h := &uploadHandler{hdlr}
	return s.Handle(s.NewHandler(&Upload{h}, opts...))
}

type uploadHandler struct {
	UploadHandler
}

func (h *uploadHandler) GetPicturePath(ctx context.Context, in *UploadPictureRequest, out *UploadPictureResponse) error {
	return h.UploadHandler.GetPicturePath(ctx, in, out)
}

func (h *uploadHandler) GetFilePath(ctx context.Context, in *UploadFileRequest, out *UploadFileResponse) error {
	return h.UploadHandler.GetFilePath(ctx, in, out)
}

func (h *uploadHandler) CheckText(ctx context.Context, in *CheckTextRequest, out *CheckTextResponse) error {
	return h.UploadHandler.CheckText(ctx, in, out)
}