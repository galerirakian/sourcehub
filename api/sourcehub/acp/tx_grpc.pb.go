// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: sourcehub/acp/tx.proto

package acp

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
	Msg_UpdateParams_FullMethodName       = "/sourcehub.acp.Msg/UpdateParams"
	Msg_CreatePolicy_FullMethodName       = "/sourcehub.acp.Msg/CreatePolicy"
	Msg_SetRelationship_FullMethodName    = "/sourcehub.acp.Msg/SetRelationship"
	Msg_DeleteRelationship_FullMethodName = "/sourcehub.acp.Msg/DeleteRelationship"
	Msg_RegisterObject_FullMethodName     = "/sourcehub.acp.Msg/RegisterObject"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error)
	SetRelationship(ctx context.Context, in *MsgSetRelationship, opts ...grpc.CallOption) (*MsgSetRelationshipResponse, error)
	DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error)
	RegisterObject(ctx context.Context, in *MsgRegisterObject, opts ...grpc.CallOption) (*MsgRegisterObjectResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreatePolicy(ctx context.Context, in *MsgCreatePolicy, opts ...grpc.CallOption) (*MsgCreatePolicyResponse, error) {
	out := new(MsgCreatePolicyResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePolicy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SetRelationship(ctx context.Context, in *MsgSetRelationship, opts ...grpc.CallOption) (*MsgSetRelationshipResponse, error) {
	out := new(MsgSetRelationshipResponse)
	err := c.cc.Invoke(ctx, Msg_SetRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteRelationship(ctx context.Context, in *MsgDeleteRelationship, opts ...grpc.CallOption) (*MsgDeleteRelationshipResponse, error) {
	out := new(MsgDeleteRelationshipResponse)
	err := c.cc.Invoke(ctx, Msg_DeleteRelationship_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RegisterObject(ctx context.Context, in *MsgRegisterObject, opts ...grpc.CallOption) (*MsgRegisterObjectResponse, error) {
	out := new(MsgRegisterObjectResponse)
	err := c.cc.Invoke(ctx, Msg_RegisterObject_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error)
	SetRelationship(context.Context, *MsgSetRelationship) (*MsgSetRelationshipResponse, error)
	DeleteRelationship(context.Context, *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error)
	RegisterObject(context.Context, *MsgRegisterObject) (*MsgRegisterObjectResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) CreatePolicy(context.Context, *MsgCreatePolicy) (*MsgCreatePolicyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolicy not implemented")
}
func (UnimplementedMsgServer) SetRelationship(context.Context, *MsgSetRelationship) (*MsgSetRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRelationship not implemented")
}
func (UnimplementedMsgServer) DeleteRelationship(context.Context, *MsgDeleteRelationship) (*MsgDeleteRelationshipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRelationship not implemented")
}
func (UnimplementedMsgServer) RegisterObject(context.Context, *MsgRegisterObject) (*MsgRegisterObjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterObject not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreatePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePolicy)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePolicy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePolicy(ctx, req.(*MsgCreatePolicy))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SetRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSetRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SetRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_SetRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SetRelationship(ctx, req.(*MsgSetRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteRelationship_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteRelationship)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteRelationship(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DeleteRelationship_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteRelationship(ctx, req.(*MsgDeleteRelationship))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RegisterObject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRegisterObject)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RegisterObject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RegisterObject_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RegisterObject(ctx, req.(*MsgRegisterObject))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sourcehub.acp.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "CreatePolicy",
			Handler:    _Msg_CreatePolicy_Handler,
		},
		{
			MethodName: "SetRelationship",
			Handler:    _Msg_SetRelationship_Handler,
		},
		{
			MethodName: "DeleteRelationship",
			Handler:    _Msg_DeleteRelationship_Handler,
		},
		{
			MethodName: "RegisterObject",
			Handler:    _Msg_RegisterObject_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sourcehub/acp/tx.proto",
}
