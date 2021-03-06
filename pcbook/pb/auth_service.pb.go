// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth_service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type LoginRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type LoginResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f39bb026ca10b68, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "techschool.pcbook.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "techschool.pcbook.LoginResponse")
}

func init() {
	proto.RegisterFile("auth_service.proto", fileDescriptor_0f39bb026ca10b68)
}

var fileDescriptor_0f39bb026ca10b68 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x49, 0x25, 0x10, 0x5c, 0xcb, 0x80, 0xa7, 0xaa, 0x4b, 0x4b, 0x16, 0x98, 0x3c, 0x94,
	0x27, 0x80, 0x81, 0xa9, 0x03, 0x0a, 0x4c, 0x30, 0x44, 0xb6, 0x39, 0xc5, 0x51, 0x12, 0x9f, 0xf1,
	0xd9, 0xf0, 0xfa, 0x28, 0x71, 0x40, 0x48, 0xc0, 0xf8, 0xfd, 0xd1, 0x77, 0x3f, 0x1d, 0x08, 0x95,
	0xa2, 0xad, 0x19, 0xc3, 0x7b, 0x6b, 0x50, 0xfa, 0x40, 0x91, 0xc4, 0x45, 0x44, 0x63, 0xd9, 0x58,
	0xa2, 0x5e, 0x7a, 0xa3, 0x89, 0xba, 0xf2, 0x1e, 0x56, 0x07, 0x6a, 0x5a, 0x57, 0xe1, 0x5b, 0x42,
	0x8e, 0x62, 0x03, 0xa7, 0x89, 0x31, 0x38, 0x35, 0xe0, 0xba, 0xd8, 0x15, 0xd7, 0x67, 0xd5, 0xb7,
	0x1e, 0x33, 0xaf, 0x98, 0x3f, 0x28, 0xbc, 0xae, 0x17, 0x39, 0xfb, 0xd2, 0xe5, 0x1e, 0xce, 0xe7,
	0x1d, 0xf6, 0xe4, 0x18, 0xc5, 0x25, 0xac, 0x94, 0x31, 0xc8, 0x5c, 0x47, 0xea, 0xd0, 0xcd, 0x63,
	0xcb, 0xec, 0x3d, 0x8d, 0xd6, 0xfe, 0x05, 0x96, 0xb7, 0x29, 0xda, 0xc7, 0xcc, 0x28, 0x0e, 0x70,
	0x3c, 0x4d, 0x88, 0xad, 0xfc, 0xc5, 0x29, 0x7f, 0x42, 0x6e, 0x76, 0xff, 0x17, 0xf2, 0xf5, 0xf2,
	0xe8, 0xee, 0x0a, 0xb6, 0x86, 0x06, 0xd9, 0xb4, 0xb1, 0x57, 0xfa, 0x8f, 0xbe, 0xd7, 0x0f, 0xc5,
	0xf3, 0xc2, 0x6b, 0x7d, 0x32, 0xfd, 0xe6, 0xe6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x1f, 0x66,
	0x22, 0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/techschool.pcbook.AuthService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
type AuthServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

// UnimplementedAuthServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (*UnimplementedAuthServiceServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}

func RegisterAuthServiceServer(s *grpc.Server, srv AuthServiceServer) {
	s.RegisterService(&_AuthService_serviceDesc, srv)
}

func _AuthService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/techschool.pcbook.AuthService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "techschool.pcbook.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _AuthService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth_service.proto",
}
