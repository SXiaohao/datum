// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: upload.proto

package upload

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UploadPictureRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName   string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileType   string `protobuf:"bytes,2,opt,name=fileType,proto3" json:"fileType,omitempty"`
	Size       int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Content    []byte `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	UploadPath string `protobuf:"bytes,5,opt,name=uploadPath,proto3" json:"uploadPath,omitempty"`
}

func (x *UploadPictureRequest) Reset() {
	*x = UploadPictureRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadPictureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPictureRequest) ProtoMessage() {}

func (x *UploadPictureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPictureRequest.ProtoReflect.Descriptor instead.
func (*UploadPictureRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{0}
}

func (x *UploadPictureRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadPictureRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UploadPictureRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadPictureRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *UploadPictureRequest) GetUploadPath() string {
	if x != nil {
		return x.UploadPath
	}
	return ""
}

type UploadPictureResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Path       string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	StatusCode int64  `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *UploadPictureResponse) Reset() {
	*x = UploadPictureResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadPictureResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadPictureResponse) ProtoMessage() {}

func (x *UploadPictureResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadPictureResponse.ProtoReflect.Descriptor instead.
func (*UploadPictureResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{1}
}

func (x *UploadPictureResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UploadPictureResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadPictureResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName string `protobuf:"bytes,1,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileType string `protobuf:"bytes,2,opt,name=fileType,proto3" json:"fileType,omitempty"`
	Size     int64  `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`
	Content  []byte `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileRequest) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRequest) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *UploadFileRequest) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *UploadFileRequest) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Path       string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	StatusCode int64  `protobuf:"varint,3,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *UploadFileResponse) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *UploadFileResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

type CheckTextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CheckTextRequest) Reset() {
	*x = CheckTextRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextRequest) ProtoMessage() {}

func (x *CheckTextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextRequest.ProtoReflect.Descriptor instead.
func (*CheckTextRequest) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{4}
}

func (x *CheckTextRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CheckTextResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg        string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	StatusCode int64  `protobuf:"varint,2,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
}

func (x *CheckTextResponse) Reset() {
	*x = CheckTextResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_upload_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckTextResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckTextResponse) ProtoMessage() {}

func (x *CheckTextResponse) ProtoReflect() protoreflect.Message {
	mi := &file_upload_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckTextResponse.ProtoReflect.Descriptor instead.
func (*CheckTextResponse) Descriptor() ([]byte, []int) {
	return file_upload_proto_rawDescGZIP(), []int{5}
}

func (x *CheckTextResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CheckTextResponse) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

var File_upload_proto protoreflect.FileDescriptor

var file_upload_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c,
	0x01, 0x0a, 0x14, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x22, 0x5d, 0x0a,
	0x15, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x22, 0x79, 0x0a, 0x11,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x5a, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x22, 0x26, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f,
	0x64, 0x65, 0x32, 0xbb, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x41, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x15, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x12, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x12, 0x11, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54,
	0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x54, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_upload_proto_rawDescOnce sync.Once
	file_upload_proto_rawDescData = file_upload_proto_rawDesc
)

func file_upload_proto_rawDescGZIP() []byte {
	file_upload_proto_rawDescOnce.Do(func() {
		file_upload_proto_rawDescData = protoimpl.X.CompressGZIP(file_upload_proto_rawDescData)
	})
	return file_upload_proto_rawDescData
}

var file_upload_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_upload_proto_goTypes = []interface{}{
	(*UploadPictureRequest)(nil),  // 0: UploadPictureRequest
	(*UploadPictureResponse)(nil), // 1: UploadPictureResponse
	(*UploadFileRequest)(nil),     // 2: UploadFileRequest
	(*UploadFileResponse)(nil),    // 3: UploadFileResponse
	(*CheckTextRequest)(nil),      // 4: CheckTextRequest
	(*CheckTextResponse)(nil),     // 5: CheckTextResponse
}
var file_upload_proto_depIdxs = []int32{
	0, // 0: Upload.GetPicturePath:input_type -> UploadPictureRequest
	2, // 1: Upload.GetFilePath:input_type -> UploadFileRequest
	4, // 2: Upload.CheckText:input_type -> CheckTextRequest
	1, // 3: Upload.GetPicturePath:output_type -> UploadPictureResponse
	3, // 4: Upload.GetFilePath:output_type -> UploadFileResponse
	5, // 5: Upload.CheckText:output_type -> CheckTextResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_upload_proto_init() }
func file_upload_proto_init() {
	if File_upload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_upload_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadPictureRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_upload_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadPictureResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_upload_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_upload_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_upload_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_upload_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckTextResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_upload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_upload_proto_goTypes,
		DependencyIndexes: file_upload_proto_depIdxs,
		MessageInfos:      file_upload_proto_msgTypes,
	}.Build()
	File_upload_proto = out.File
	file_upload_proto_rawDesc = nil
	file_upload_proto_goTypes = nil
	file_upload_proto_depIdxs = nil
}
