// Code generated by protoc-gen-go. DO NOT EDIT.
// source: iam.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type SignInOrSignUpRequest struct {
	NickName             *wrappers.StringValue `protobuf:"bytes,1,opt,name=nick_name,json=nickName,proto3" json:"nick_name,omitempty"`
	Password             *wrappers.StringValue `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SignInOrSignUpRequest) Reset()         { *m = SignInOrSignUpRequest{} }
func (m *SignInOrSignUpRequest) String() string { return proto.CompactTextString(m) }
func (*SignInOrSignUpRequest) ProtoMessage()    {}
func (*SignInOrSignUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{0}
}

func (m *SignInOrSignUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInOrSignUpRequest.Unmarshal(m, b)
}
func (m *SignInOrSignUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInOrSignUpRequest.Marshal(b, m, deterministic)
}
func (m *SignInOrSignUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInOrSignUpRequest.Merge(m, src)
}
func (m *SignInOrSignUpRequest) XXX_Size() int {
	return xxx_messageInfo_SignInOrSignUpRequest.Size(m)
}
func (m *SignInOrSignUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInOrSignUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInOrSignUpRequest proto.InternalMessageInfo

func (m *SignInOrSignUpRequest) GetNickName() *wrappers.StringValue {
	if m != nil {
		return m.NickName
	}
	return nil
}

func (m *SignInOrSignUpRequest) GetPassword() *wrappers.StringValue {
	if m != nil {
		return m.Password
	}
	return nil
}

type SignInOrSignUpResponse struct {
	AccessToken          *wrappers.StringValue `protobuf:"bytes,3,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *SignInOrSignUpResponse) Reset()         { *m = SignInOrSignUpResponse{} }
func (m *SignInOrSignUpResponse) String() string { return proto.CompactTextString(m) }
func (*SignInOrSignUpResponse) ProtoMessage()    {}
func (*SignInOrSignUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{1}
}

func (m *SignInOrSignUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInOrSignUpResponse.Unmarshal(m, b)
}
func (m *SignInOrSignUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInOrSignUpResponse.Marshal(b, m, deterministic)
}
func (m *SignInOrSignUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInOrSignUpResponse.Merge(m, src)
}
func (m *SignInOrSignUpResponse) XXX_Size() int {
	return xxx_messageInfo_SignInOrSignUpResponse.Size(m)
}
func (m *SignInOrSignUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInOrSignUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignInOrSignUpResponse proto.InternalMessageInfo

func (m *SignInOrSignUpResponse) GetAccessToken() *wrappers.StringValue {
	if m != nil {
		return m.AccessToken
	}
	return nil
}

type SignOutRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutRequest) Reset()         { *m = SignOutRequest{} }
func (m *SignOutRequest) String() string { return proto.CompactTextString(m) }
func (*SignOutRequest) ProtoMessage()    {}
func (*SignOutRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{2}
}

func (m *SignOutRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutRequest.Unmarshal(m, b)
}
func (m *SignOutRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutRequest.Marshal(b, m, deterministic)
}
func (m *SignOutRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutRequest.Merge(m, src)
}
func (m *SignOutRequest) XXX_Size() int {
	return xxx_messageInfo_SignOutRequest.Size(m)
}
func (m *SignOutRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutRequest proto.InternalMessageInfo

type SignOutResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignOutResponse) Reset()         { *m = SignOutResponse{} }
func (m *SignOutResponse) String() string { return proto.CompactTextString(m) }
func (*SignOutResponse) ProtoMessage()    {}
func (*SignOutResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2c201915207782, []int{3}
}

func (m *SignOutResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignOutResponse.Unmarshal(m, b)
}
func (m *SignOutResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignOutResponse.Marshal(b, m, deterministic)
}
func (m *SignOutResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignOutResponse.Merge(m, src)
}
func (m *SignOutResponse) XXX_Size() int {
	return xxx_messageInfo_SignOutResponse.Size(m)
}
func (m *SignOutResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SignOutResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SignOutResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SignInOrSignUpRequest)(nil), "ltk.SignInOrSignUpRequest")
	proto.RegisterType((*SignInOrSignUpResponse)(nil), "ltk.SignInOrSignUpResponse")
	proto.RegisterType((*SignOutRequest)(nil), "ltk.SignOutRequest")
	proto.RegisterType((*SignOutResponse)(nil), "ltk.SignOutResponse")
}

func init() { proto.RegisterFile("iam.proto", fileDescriptor_0a2c201915207782) }

var fileDescriptor_0a2c201915207782 = []byte{
	// 409 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcc, 0x4c, 0xcc, 0xd5,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x29, 0xc9, 0x96, 0x92, 0x4b, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x17, 0x25, 0x16, 0x14, 0xa4, 0x16,
	0x15, 0x43, 0x14, 0x49, 0xc9, 0x40, 0xe5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x60, 0xb2, 0x3a, 0x60, 0x2a, 0x59, 0x37, 0x3d, 0x35, 0x4f, 0xb7,
	0xb8, 0x3c, 0x31, 0x3d, 0x3d, 0xb5, 0x48, 0x3f, 0xbf, 0x00, 0xac, 0x02, 0x53, 0xb5, 0x52, 0x0f,
	0x23, 0x97, 0x68, 0x70, 0x66, 0x7a, 0x9e, 0x67, 0x9e, 0x7f, 0x11, 0x88, 0x0e, 0x2d, 0x08, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xb2, 0xe4, 0xe2, 0xcc, 0xcb, 0x4c, 0xce, 0x8e, 0xcf, 0x4b,
	0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x92, 0xd1, 0x83, 0xd8, 0xac, 0x07, 0x73,
	0x99, 0x5e, 0x70, 0x49, 0x51, 0x66, 0x5e, 0x7a, 0x58, 0x62, 0x4e, 0x69, 0x6a, 0x10, 0x07, 0x48,
	0xb9, 0x5f, 0x62, 0x6e, 0xaa, 0x90, 0x05, 0x17, 0x47, 0x41, 0x62, 0x71, 0x71, 0x79, 0x7e, 0x51,
	0x8a, 0x04, 0x13, 0x31, 0x3a, 0x61, 0xaa, 0x95, 0x22, 0xb9, 0xc4, 0xd0, 0x5d, 0x53, 0x5c, 0x90,
	0x9f, 0x57, 0x9c, 0x2a, 0x64, 0xcf, 0xc5, 0x93, 0x98, 0x9c, 0x9c, 0x5a, 0x5c, 0x1c, 0x5f, 0x92,
	0x9f, 0x9d, 0x9a, 0x27, 0xc1, 0x4c, 0x84, 0xb9, 0xdc, 0x10, 0x1d, 0x21, 0x20, 0x0d, 0x4a, 0x02,
	0x5c, 0x7c, 0x20, 0x23, 0xfd, 0x4b, 0x4b, 0xa0, 0x3e, 0x54, 0x12, 0xe4, 0xe2, 0x87, 0x8b, 0x40,
	0x6c, 0x31, 0xfa, 0xcd, 0xc8, 0xc5, 0x03, 0x56, 0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a,
	0xd4, 0xc5, 0x08, 0xd1, 0x86, 0x70, 0x91, 0x90, 0x94, 0x5e, 0x4e, 0x49, 0xb6, 0x1e, 0xd6, 0x40,
	0x93, 0x92, 0xc6, 0x2a, 0x07, 0x31, 0x5c, 0xc9, 0x7e, 0x92, 0x23, 0x9f, 0x10, 0xcf, 0xd3, 0xfd,
	0xab, 0x5f, 0x36, 0xcc, 0x7f, 0x3e, 0x73, 0xf7, 0xd3, 0xbd, 0x53, 0x9b, 0x2e, 0x3f, 0x99, 0xcc,
	0xa4, 0xa1, 0xa4, 0xac, 0x5f, 0x66, 0xa8, 0x9f, 0x9f, 0x58, 0x5a, 0x92, 0x61, 0xa4, 0x0f, 0xf6,
	0x9d, 0x55, 0x71, 0x66, 0x7a, 0x5e, 0x7c, 0x66, 0x5e, 0x7c, 0x7e, 0x51, 0x3c, 0x98, 0x59, 0x5a,
	0x60, 0xc5, 0xa8, 0x25, 0x14, 0xc7, 0xc5, 0x0e, 0x75, 0xb0, 0x90, 0x30, 0xdc, 0x22, 0x84, 0x87,
	0xa4, 0x44, 0x50, 0x05, 0xa1, 0xd6, 0x6a, 0x4c, 0x72, 0xe4, 0x10, 0x62, 0x03, 0x59, 0xd8, 0xbe,
	0x0b, 0x6c, 0xa1, 0xb8, 0x92, 0x10, 0xc8, 0x42, 0x24, 0x9b, 0xf2, 0x4b, 0x4b, 0xac, 0x18, 0xb5,
	0x9c, 0x58, 0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x61, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0x50, 0xb3, 0xae, 0x78, 0x97, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TokenServiceClient interface {
	SignInOrSignUp(ctx context.Context, in *SignInOrSignUpRequest, opts ...grpc.CallOption) (*SignInOrSignUpResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
}

type tokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewTokenServiceClient(cc *grpc.ClientConn) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) SignInOrSignUp(ctx context.Context, in *SignInOrSignUpRequest, opts ...grpc.CallOption) (*SignInOrSignUpResponse, error) {
	out := new(SignInOrSignUpResponse)
	err := c.cc.Invoke(ctx, "/ltk.TokenService/SignInOrSignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/ltk.TokenService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenServiceServer is the server API for TokenService service.
type TokenServiceServer interface {
	SignInOrSignUp(context.Context, *SignInOrSignUpRequest) (*SignInOrSignUpResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
}

// UnimplementedTokenServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (*UnimplementedTokenServiceServer) SignInOrSignUp(ctx context.Context, req *SignInOrSignUpRequest) (*SignInOrSignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignInOrSignUp not implemented")
}
func (*UnimplementedTokenServiceServer) SignOut(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}

func RegisterTokenServiceServer(s *grpc.Server, srv TokenServiceServer) {
	s.RegisterService(&_TokenService_serviceDesc, srv)
}

func _TokenService_SignInOrSignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInOrSignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).SignInOrSignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ltk.TokenService/SignInOrSignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).SignInOrSignUp(ctx, req.(*SignInOrSignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ltk.TokenService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ltk.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignInOrSignUp",
			Handler:    _TokenService_SignInOrSignUp_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _TokenService_SignOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "iam.proto",
}
