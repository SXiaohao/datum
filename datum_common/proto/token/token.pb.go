// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.4
// source: token.proto

package proto

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

type GetTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *GetTokenRequest) Reset() {
	*x = GetTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenRequest) ProtoMessage() {}

func (x *GetTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenRequest.ProtoReflect.Descriptor instead.
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{0}
}

func (x *GetTokenRequest) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Msg   string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetTokenResponse) Reset() {
	*x = GetTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenResponse) ProtoMessage() {}

func (x *GetTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenResponse.ProtoReflect.Descriptor instead.
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *GetTokenResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ParseTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ParseTokenRequest) Reset() {
	*x = ParseTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenRequest) ProtoMessage() {}

func (x *ParseTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenRequest.ProtoReflect.Descriptor instead.
func (*ParseTokenRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{2}
}

func (x *ParseTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ParseTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid uint32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ParseTokenResponse) Reset() {
	*x = ParseTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ParseTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParseTokenResponse) ProtoMessage() {}

func (x *ParseTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParseTokenResponse.ProtoReflect.Descriptor instead.
func (*ParseTokenResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{3}
}

func (x *ParseTokenResponse) GetUid() uint32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *ParseTokenResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_token_proto protoreflect.FileDescriptor

var file_token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x29, 0x0a, 0x11, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x38, 0x0a, 0x12, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x8e, 0x01, 0x0a, 0x0c, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x50, 0x61, 0x72, 0x73,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x61, 0x72, 0x73, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x61, 0x72, 0x73, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_token_proto_rawDescOnce sync.Once
	file_token_proto_rawDescData = file_token_proto_rawDesc
)

func file_token_proto_rawDescGZIP() []byte {
	file_token_proto_rawDescOnce.Do(func() {
		file_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_proto_rawDescData)
	})
	return file_token_proto_rawDescData
}

var file_token_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_token_proto_goTypes = []interface{}{
	(*GetTokenRequest)(nil),    // 0: proto.GetTokenRequest
	(*GetTokenResponse)(nil),   // 1: proto.GetTokenResponse
	(*ParseTokenRequest)(nil),  // 2: proto.ParseTokenRequest
	(*ParseTokenResponse)(nil), // 3: proto.ParseTokenResponse
}
var file_token_proto_depIdxs = []int32{
	0, // 0: proto.TokenService.GetToken:input_type -> proto.GetTokenRequest
	2, // 1: proto.TokenService.ParseToken:input_type -> proto.ParseTokenRequest
	1, // 2: proto.TokenService.GetToken:output_type -> proto.GetTokenResponse
	3, // 3: proto.TokenService.ParseToken:output_type -> proto.ParseTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_proto_init() }
func file_token_proto_init() {
	if File_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenRequest); i {
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
		file_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTokenResponse); i {
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
		file_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenRequest); i {
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
		file_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ParseTokenResponse); i {
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
			RawDescriptor: file_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_proto_goTypes,
		DependencyIndexes: file_token_proto_depIdxs,
		MessageInfos:      file_token_proto_msgTypes,
	}.Build()
	File_token_proto = out.File
	file_token_proto_rawDesc = nil
	file_token_proto_goTypes = nil
	file_token_proto_depIdxs = nil
}
