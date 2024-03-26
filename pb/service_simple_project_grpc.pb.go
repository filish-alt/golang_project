// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.0
// source: service_simple_project.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Simpleproject_Createuser_FullMethodName = "/pb.simpleproject/Createuser"
	Simpleproject_Verifyuser_FullMethodName = "/pb.simpleproject/Verifyuser"
	Simpleproject_Updateuser_FullMethodName = "/pb.simpleproject/Updateuser"
	Simpleproject_Loginuser_FullMethodName  = "/pb.simpleproject/Loginuser"
)

// SimpleprojectClient is the client API for Simpleproject service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SimpleprojectClient interface {
	Createuser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	Verifyuser(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error)
	Updateuser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	Loginuser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type simpleprojectClient struct {
	cc grpc.ClientConnInterface
}

func NewSimpleprojectClient(cc grpc.ClientConnInterface) SimpleprojectClient {
	return &simpleprojectClient{cc}
}

func (c *simpleprojectClient) Createuser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Simpleproject_Createuser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleprojectClient) Verifyuser(ctx context.Context, in *VerifyEmailRequest, opts ...grpc.CallOption) (*VerifyEmailResponse, error) {
	out := new(VerifyEmailResponse)
	err := c.cc.Invoke(ctx, Simpleproject_Verifyuser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleprojectClient) Updateuser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, Simpleproject_Updateuser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleprojectClient) Loginuser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, Simpleproject_Loginuser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SimpleprojectServer is the server API for Simpleproject service.
// All implementations must embed UnimplementedSimpleprojectServer
// for forward compatibility
type SimpleprojectServer interface {
	Createuser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	Verifyuser(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error)
	Updateuser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	Loginuser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
	mustEmbedUnimplementedSimpleprojectServer()
}

// UnimplementedSimpleprojectServer must be embedded to have forward compatible implementations.
type UnimplementedSimpleprojectServer struct {
}

func (UnimplementedSimpleprojectServer) Createuser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Createuser not implemented")
}
func (UnimplementedSimpleprojectServer) Verifyuser(context.Context, *VerifyEmailRequest) (*VerifyEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Verifyuser not implemented")
}
func (UnimplementedSimpleprojectServer) Updateuser(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Updateuser not implemented")
}
func (UnimplementedSimpleprojectServer) Loginuser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Loginuser not implemented")
}
func (UnimplementedSimpleprojectServer) mustEmbedUnimplementedSimpleprojectServer() {}

// UnsafeSimpleprojectServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SimpleprojectServer will
// result in compilation errors.
type UnsafeSimpleprojectServer interface {
	mustEmbedUnimplementedSimpleprojectServer()
}

func RegisterSimpleprojectServer(s grpc.ServiceRegistrar, srv SimpleprojectServer) {
	s.RegisterService(&Simpleproject_ServiceDesc, srv)
}

func _Simpleproject_Createuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleprojectServer).Createuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Simpleproject_Createuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleprojectServer).Createuser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simpleproject_Verifyuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleprojectServer).Verifyuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Simpleproject_Verifyuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleprojectServer).Verifyuser(ctx, req.(*VerifyEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simpleproject_Updateuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleprojectServer).Updateuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Simpleproject_Updateuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleprojectServer).Updateuser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Simpleproject_Loginuser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleprojectServer).Loginuser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Simpleproject_Loginuser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleprojectServer).Loginuser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Simpleproject_ServiceDesc is the grpc.ServiceDesc for Simpleproject service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Simpleproject_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.simpleproject",
	HandlerType: (*SimpleprojectServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Createuser",
			Handler:    _Simpleproject_Createuser_Handler,
		},
		{
			MethodName: "Verifyuser",
			Handler:    _Simpleproject_Verifyuser_Handler,
		},
		{
			MethodName: "Updateuser",
			Handler:    _Simpleproject_Updateuser_Handler,
		},
		{
			MethodName: "Loginuser",
			Handler:    _Simpleproject_Loginuser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_simple_project.proto",
}
