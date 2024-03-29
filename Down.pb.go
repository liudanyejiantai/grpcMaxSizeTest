// Code generated by protoc-gen-go.
// source: Down.proto
// DO NOT EDIT!

/*
Package PbDownMaxSize is a generated protocol buffer package.

It is generated from these files:
	Down.proto

It has these top-level messages:
	ReqDown
	ResDown
*/
package PbDownMaxSize

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 下载的请求参数
type ReqDown struct {
	// 文件名称
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *ReqDown) Reset()                    { *m = ReqDown{} }
func (m *ReqDown) String() string            { return proto.CompactTextString(m) }
func (*ReqDown) ProtoMessage()               {}
func (*ReqDown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ReqDown) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// 应答消息
type ResDown struct {
	// 返回提示消息
	RetMsg string `protobuf:"bytes,1,opt,name=retMsg" json:"retMsg,omitempty"`
	// 返回的碎片影像的buffer
	FileBuf []byte `protobuf:"bytes,2,opt,name=fileBuf,proto3" json:"fileBuf,omitempty"`
	// 返回的状态码(0表示成功,其他失败)
	IRet int32 `protobuf:"varint,3,opt,name=iRet" json:"iRet,omitempty"`
}

func (m *ResDown) Reset()                    { *m = ResDown{} }
func (m *ResDown) String() string            { return proto.CompactTextString(m) }
func (*ResDown) ProtoMessage()               {}
func (*ResDown) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ResDown) GetRetMsg() string {
	if m != nil {
		return m.RetMsg
	}
	return ""
}

func (m *ResDown) GetFileBuf() []byte {
	if m != nil {
		return m.FileBuf
	}
	return nil
}

func (m *ResDown) GetIRet() int32 {
	if m != nil {
		return m.IRet
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqDown)(nil), "PbDownMaxSize.reqDown")
	proto.RegisterType((*ResDown)(nil), "PbDownMaxSize.resDown")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PbDownMaxSize service

type PbDownMaxSizeClient interface {
	// RPC上传大文件分片
	DownLoad(ctx context.Context, in *ReqDown, opts ...grpc.CallOption) (*ResDown, error)
}

type pbDownMaxSizeClient struct {
	cc *grpc.ClientConn
}

func NewPbDownMaxSizeClient(cc *grpc.ClientConn) PbDownMaxSizeClient {
	return &pbDownMaxSizeClient{cc}
}

func (c *pbDownMaxSizeClient) DownLoad(ctx context.Context, in *ReqDown, opts ...grpc.CallOption) (*ResDown, error) {
	out := new(ResDown)
	err := grpc.Invoke(ctx, "/PbDownMaxSize.PbDownMaxSize/DownLoad", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PbDownMaxSize service

type PbDownMaxSizeServer interface {
	// RPC上传大文件分片
	DownLoad(context.Context, *ReqDown) (*ResDown, error)
}

func RegisterPbDownMaxSizeServer(s *grpc.Server, srv PbDownMaxSizeServer) {
	s.RegisterService(&_PbDownMaxSize_serviceDesc, srv)
}

func _PbDownMaxSize_DownLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqDown)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PbDownMaxSizeServer).DownLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/PbDownMaxSize.PbDownMaxSize/DownLoad",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PbDownMaxSizeServer).DownLoad(ctx, req.(*ReqDown))
	}
	return interceptor(ctx, in, info, handler)
}

var _PbDownMaxSize_serviceDesc = grpc.ServiceDesc{
	ServiceName: "PbDownMaxSize.PbDownMaxSize",
	HandlerType: (*PbDownMaxSizeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownLoad",
			Handler:    _PbDownMaxSize_DownLoad_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Down.proto",
}

func init() { proto.RegisterFile("Down.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x72, 0xc9, 0x2f, 0xcf,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x0d, 0x48, 0x02, 0xf1, 0x7c, 0x13, 0x2b, 0x82,
	0x33, 0xab, 0x52, 0x95, 0x64, 0xb9, 0xd8, 0x8b, 0x52, 0x0b, 0x41, 0x22, 0x42, 0x42, 0x5c, 0x2c,
	0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x60, 0xb6, 0x92, 0x3f, 0x48,
	0xba, 0x18, 0x2c, 0x2d, 0xc6, 0xc5, 0x56, 0x94, 0x5a, 0xe2, 0x5b, 0x9c, 0x0e, 0x55, 0x00, 0xe5,
	0x09, 0x49, 0x70, 0xb1, 0xa7, 0x65, 0xe6, 0xa4, 0x3a, 0x95, 0xa6, 0x49, 0x30, 0x01, 0x25, 0x78,
	0x82, 0x60, 0x5c, 0x90, 0x81, 0x99, 0x41, 0xa9, 0x25, 0x12, 0xcc, 0x40, 0x61, 0xd6, 0x20, 0x30,
	0xdb, 0xc8, 0x97, 0x0b, 0xd5, 0x01, 0x42, 0x36, 0x5c, 0x1c, 0x20, 0xae, 0x4f, 0x7e, 0x62, 0x8a,
	0x90, 0x98, 0x1e, 0x8a, 0x9c, 0x1e, 0xd4, 0x65, 0x52, 0x98, 0xe2, 0x60, 0x27, 0x29, 0x31, 0x24,
	0xb1, 0x81, 0x3d, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x9c, 0x00, 0x98, 0xeb, 0xe2, 0x00,
	0x00, 0x00,
}
