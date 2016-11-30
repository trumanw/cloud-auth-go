// Code generated by protoc-gen-go.
// source: pb/token.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Token struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	TokenType    string `protobuf:"bytes,2,opt,name=token_type,json=tokenType" json:"token_type,omitempty"`
	ExpiresIn    uint64 `protobuf:"varint,3,opt,name=expires_in,json=expiresIn" json:"expires_in,omitempty"`
	RefreshToken string `protobuf:"bytes,4,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Scope        string `protobuf:"bytes,5,opt,name=scope" json:"scope,omitempty"`
	State        string `protobuf:"bytes,6,opt,name=state" json:"state,omitempty"`
}

func (m *Token) Reset()                    { *m = Token{} }
func (m *Token) String() string            { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()               {}
func (*Token) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Token) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Token) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *Token) GetExpiresIn() uint64 {
	if m != nil {
		return m.ExpiresIn
	}
	return 0
}

func (m *Token) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Token) GetScope() string {
	if m != nil {
		return m.Scope
	}
	return ""
}

func (m *Token) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Token)(nil), "pb.Token")
}

func init() { proto.RegisterFile("pb/token.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x2c, 0xcf, 0xb1, 0xca, 0xc2, 0x30,
	0x14, 0x05, 0x60, 0xd2, 0xbf, 0x2d, 0xf4, 0xfe, 0xd5, 0xa1, 0x38, 0x64, 0x11, 0xaa, 0x2e, 0x9d,
	0x74, 0xf0, 0x29, 0x5c, 0x4b, 0xf7, 0xd2, 0x84, 0x2b, 0x16, 0x21, 0xb9, 0x24, 0x19, 0xec, 0xa3,
	0xf9, 0x76, 0x92, 0x9b, 0x8c, 0xf7, 0x3b, 0x87, 0x1c, 0x02, 0x7b, 0x52, 0xb7, 0x60, 0xdf, 0x68,
	0xae, 0xe4, 0x6c, 0xb0, 0x5d, 0x41, 0xea, 0xfc, 0x15, 0x50, 0x4d, 0xd1, 0xba, 0x13, 0xb4, 0x8b,
	0xd6, 0xe8, 0xfd, 0xcc, 0x1d, 0x29, 0x7a, 0x31, 0x34, 0xe3, 0x7f, 0xb2, 0x54, 0x39, 0x02, 0x70,
	0x36, 0x87, 0x8d, 0x50, 0x16, 0x5c, 0x68, 0x58, 0xa6, 0x8d, 0x30, 0xc6, 0xf8, 0xa1, 0xd5, 0xa1,
	0x9f, 0x57, 0x23, 0xff, 0x7a, 0x31, 0x94, 0x63, 0x93, 0xe5, 0x61, 0xba, 0x0b, 0xec, 0x1c, 0x3e,
	0x1d, 0xfa, 0x57, 0x5e, 0x28, 0xf9, 0x81, 0x36, 0x63, 0x9a, 0x38, 0x40, 0xe5, 0xb5, 0x25, 0x94,
	0x15, 0x87, 0xe9, 0x60, 0x0d, 0x4b, 0x40, 0x59, 0x67, 0x8d, 0x87, 0xaa, 0xf9, 0x1b, 0xf7, 0x5f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0xef, 0x7c, 0x0d, 0xd8, 0x00, 0x00, 0x00,
}
