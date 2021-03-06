// Code generated by protoc-gen-go.
// source: pb/client_credentials.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/client_credentials.proto
	pb/token.proto

It has these top-level messages:
	ClientCredentialsCreatedEvent
	Token
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

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

type ClientCredentialsCreatedEvent struct {
	ClientId      string `protobuf:"bytes,1,opt,name=client_id,json=clientId" json:"client_id,omitempty"`
	ClientSecret  string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret" json:"client_secret,omitempty"`
	RequestId     string `protobuf:"bytes,3,opt,name=request_id,json=requestId" json:"request_id,omitempty"`
	PartnerAttrId string `protobuf:"bytes,4,opt,name=partner_attr_id,json=partnerAttrId" json:"partner_attr_id,omitempty"`
	GrantType     string `protobuf:"bytes,5,opt,name=grant_type,json=grantType" json:"grant_type,omitempty"`
}

func (m *ClientCredentialsCreatedEvent) Reset()                    { *m = ClientCredentialsCreatedEvent{} }
func (m *ClientCredentialsCreatedEvent) String() string            { return proto.CompactTextString(m) }
func (*ClientCredentialsCreatedEvent) ProtoMessage()               {}
func (*ClientCredentialsCreatedEvent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ClientCredentialsCreatedEvent) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *ClientCredentialsCreatedEvent) GetClientSecret() string {
	if m != nil {
		return m.ClientSecret
	}
	return ""
}

func (m *ClientCredentialsCreatedEvent) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ClientCredentialsCreatedEvent) GetPartnerAttrId() string {
	if m != nil {
		return m.PartnerAttrId
	}
	return ""
}

func (m *ClientCredentialsCreatedEvent) GetGrantType() string {
	if m != nil {
		return m.GrantType
	}
	return ""
}

func init() {
	proto.RegisterType((*ClientCredentialsCreatedEvent)(nil), "pb.ClientCredentialsCreatedEvent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CilentCredentialsService service

type CilentCredentialsServiceClient interface {
	CreateClientCredentials(ctx context.Context, in *ClientCredentialsCreatedEvent, opts ...grpc.CallOption) (*Token, error)
}

type cilentCredentialsServiceClient struct {
	cc *grpc.ClientConn
}

func NewCilentCredentialsServiceClient(cc *grpc.ClientConn) CilentCredentialsServiceClient {
	return &cilentCredentialsServiceClient{cc}
}

func (c *cilentCredentialsServiceClient) CreateClientCredentials(ctx context.Context, in *ClientCredentialsCreatedEvent, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/pb.CilentCredentialsService/CreateClientCredentials", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CilentCredentialsService service

type CilentCredentialsServiceServer interface {
	CreateClientCredentials(context.Context, *ClientCredentialsCreatedEvent) (*Token, error)
}

func RegisterCilentCredentialsServiceServer(s *grpc.Server, srv CilentCredentialsServiceServer) {
	s.RegisterService(&_CilentCredentialsService_serviceDesc, srv)
}

func _CilentCredentialsService_CreateClientCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientCredentialsCreatedEvent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CilentCredentialsServiceServer).CreateClientCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CilentCredentialsService/CreateClientCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CilentCredentialsServiceServer).CreateClientCredentials(ctx, req.(*ClientCredentialsCreatedEvent))
	}
	return interceptor(ctx, in, info, handler)
}

var _CilentCredentialsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CilentCredentialsService",
	HandlerType: (*CilentCredentialsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClientCredentials",
			Handler:    _CilentCredentialsService_CreateClientCredentials_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/client_credentials.proto",
}

func init() { proto.RegisterFile("pb/client_credentials.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 285 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x85, 0x69, 0xfd, 0xc1, 0x06, 0xab, 0x90, 0x8d, 0xa5, 0xe3, 0x80, 0x56, 0x10, 0x71, 0xd1,
	0xc2, 0xb8, 0x73, 0x27, 0xc5, 0x45, 0xb7, 0xce, 0xec, 0x4b, 0xda, 0x5c, 0x6a, 0xb0, 0x24, 0xf1,
	0xf6, 0xce, 0xc0, 0x2c, 0xdc, 0xf8, 0x0a, 0x3e, 0x94, 0x0f, 0xe0, 0x2b, 0xf8, 0x20, 0xd2, 0xa4,
	0x20, 0x22, 0xcc, 0x32, 0x5f, 0x3e, 0xce, 0xe5, 0x1c, 0x36, 0xb3, 0x4d, 0xd1, 0xf6, 0x0a, 0x34,
	0xd5, 0x2d, 0x82, 0x04, 0x4d, 0x4a, 0xf4, 0x43, 0x6e, 0xd1, 0x90, 0xe1, 0xa1, 0x6d, 0xd2, 0xf3,
	0xce, 0x98, 0xae, 0x87, 0x42, 0x58, 0x55, 0x08, 0xad, 0x0d, 0x09, 0x52, 0x46, 0x4f, 0x46, 0x7a,
	0x62, 0x9b, 0x82, 0xcc, 0x0b, 0x68, 0xff, 0xce, 0x3e, 0x03, 0x36, 0x2f, 0x5d, 0x5c, 0xf9, 0x9b,
	0x56, 0x22, 0x08, 0x02, 0xf9, 0xb8, 0x01, 0x4d, 0x7c, 0xc6, 0xa2, 0xe9, 0x9e, 0x92, 0x49, 0x70,
	0x11, 0xdc, 0x44, 0x4f, 0x47, 0x1e, 0x54, 0x92, 0x5f, 0xb1, 0x78, 0xfa, 0x1c, 0xa0, 0x45, 0xa0,
	0x24, 0x74, 0xc2, 0xb1, 0x87, 0x4b, 0xc7, 0xf8, 0x9c, 0x31, 0x84, 0xd7, 0x35, 0x0c, 0x2e, 0x62,
	0xcf, 0x19, 0xd1, 0x44, 0x2a, 0xc9, 0xaf, 0xd9, 0xa9, 0x15, 0x48, 0x1a, 0xb0, 0x16, 0x44, 0x38,
	0x3a, 0xfb, 0xce, 0x89, 0x27, 0xfc, 0x40, 0x84, 0x95, 0x1c, 0x63, 0x3a, 0x14, 0x9a, 0x6a, 0xda,
	0x5a, 0x48, 0x0e, 0x7c, 0x8c, 0x23, 0xab, 0xad, 0x85, 0xc5, 0x1b, 0x4b, 0x4a, 0xd5, 0xff, 0x2d,
	0xb2, 0x04, 0xdc, 0xa8, 0x16, 0xb8, 0x60, 0x67, 0xbe, 0xd3, 0xbf, 0xaa, 0xfc, 0x32, 0xb7, 0x4d,
	0xbe, 0x73, 0x81, 0x34, 0x1a, 0x95, 0xd5, 0x38, 0x5a, 0x96, 0xbc, 0x7f, 0x7d, 0x7f, 0x84, 0x3c,
	0x8b, 0x0b, 0x23, 0xd6, 0xf4, 0xbc, 0xf0, 0x5b, 0xde, 0x07, 0xb7, 0xcd, 0xa1, 0xdb, 0xf3, 0xee,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x87, 0x3d, 0x91, 0x11, 0xa0, 0x01, 0x00, 0x00,
}
