// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accountspb

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

// AccountsServiceClient is the client API for AccountsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountsServiceClient interface {
	Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error)
	SetCredentials(ctx context.Context, in *SetCredentialsRequest, opts ...grpc.CallOption) (*SetCredentialsResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Account, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type accountsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountsServiceClient(cc grpc.ClientConnInterface) AccountsServiceClient {
	return &accountsServiceClient{cc}
}

func (c *accountsServiceClient) Token(ctx context.Context, in *TokenRequest, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := c.cc.Invoke(ctx, "/nocloud.accounts.AccountsService/Token", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) SetCredentials(ctx context.Context, in *SetCredentialsRequest, opts ...grpc.CallOption) (*SetCredentialsResponse, error) {
	out := new(SetCredentialsResponse)
	err := c.cc.Invoke(ctx, "/nocloud.accounts.AccountsService/SetCredentials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/nocloud.accounts.AccountsService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/nocloud.accounts.AccountsService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountsServiceClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/nocloud.accounts.AccountsService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountsServiceServer is the server API for AccountsService service.
// All implementations must embed UnimplementedAccountsServiceServer
// for forward compatibility
type AccountsServiceServer interface {
	Token(context.Context, *TokenRequest) (*TokenResponse, error)
	SetCredentials(context.Context, *SetCredentialsRequest) (*SetCredentialsResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*Account, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	mustEmbedUnimplementedAccountsServiceServer()
}

// UnimplementedAccountsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccountsServiceServer struct {
}

func (UnimplementedAccountsServiceServer) Token(context.Context, *TokenRequest) (*TokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Token not implemented")
}
func (UnimplementedAccountsServiceServer) SetCredentials(context.Context, *SetCredentialsRequest) (*SetCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetCredentials not implemented")
}
func (UnimplementedAccountsServiceServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAccountsServiceServer) Get(context.Context, *GetRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccountsServiceServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAccountsServiceServer) mustEmbedUnimplementedAccountsServiceServer() {}

// UnsafeAccountsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountsServiceServer will
// result in compilation errors.
type UnsafeAccountsServiceServer interface {
	mustEmbedUnimplementedAccountsServiceServer()
}

func RegisterAccountsServiceServer(s grpc.ServiceRegistrar, srv AccountsServiceServer) {
	s.RegisterService(&AccountsService_ServiceDesc, srv)
}

func _AccountsService_Token_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Token(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.accounts.AccountsService/Token",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Token(ctx, req.(*TokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_SetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).SetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.accounts.AccountsService/SetCredentials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).SetCredentials(ctx, req.(*SetCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.accounts.AccountsService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.accounts.AccountsService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountsService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nocloud.accounts.AccountsService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServiceServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountsService_ServiceDesc is the grpc.ServiceDesc for AccountsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nocloud.accounts.AccountsService",
	HandlerType: (*AccountsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Token",
			Handler:    _AccountsService_Token_Handler,
		},
		{
			MethodName: "SetCredentials",
			Handler:    _AccountsService_SetCredentials_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _AccountsService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AccountsService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AccountsService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/accounts/accountspb/accounts.proto",
}
