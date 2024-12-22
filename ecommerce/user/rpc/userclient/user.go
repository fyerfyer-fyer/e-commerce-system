// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package userclient

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddAddressRequest        = user.AddAddressRequest
	AddAddressResponse       = user.AddAddressResponse
	AddCollectionRequest     = user.AddCollectionRequest
	Address                  = user.Address
	DeleteAddressRequest     = user.DeleteAddressRequest
	DeleteCollectionRequest  = user.DeleteCollectionRequest
	EditAddressRequest       = user.EditAddressRequest
	EditCollectionRequest    = user.EditCollectionRequest
	Empty                    = user.Empty
	GetAddressesRequest      = user.GetAddressesRequest
	GetAddressesResponse     = user.GetAddressesResponse
	GetCollectionRequest     = user.GetCollectionRequest
	GetCollectionResponse    = user.GetCollectionResponse
	GetLoginHistoryRequest   = user.GetLoginHistoryRequest
	GetLoginHistoryResponse  = user.GetLoginHistoryResponse
	GetUserInfoRequest       = user.GetUserInfoRequest
	GetUserInfoResponse      = user.GetUserInfoResponse
	LoginHistoryEntry        = user.LoginHistoryEntry
	LoginRequest             = user.LoginRequest
	LoginResponse            = user.LoginResponse
	PaginationRequest        = user.PaginationRequest
	PaginationResponse       = user.PaginationResponse
	RegisterRequest          = user.RegisterRequest
	RegisterResponse         = user.RegisterResponse
	SetDefaultAddressRequest = user.SetDefaultAddressRequest
	UserInfo                 = user.UserInfo

	User interface {
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		GetAddresses(ctx context.Context, in *GetAddressesRequest, opts ...grpc.CallOption) (*GetAddressesResponse, error)
		AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*AddAddressResponse, error)
		EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*Empty, error)
		DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*Empty, error)
		SetDefaultAddress(ctx context.Context, in *SetDefaultAddressRequest, opts ...grpc.CallOption) (*Empty, error)
		GetLoginHistory(ctx context.Context, in *GetLoginHistoryRequest, opts ...grpc.CallOption) (*GetLoginHistoryResponse, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error)
		GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error)
		AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*Empty, error)
		EditCollection(ctx context.Context, in *EditCollectionRequest, opts ...grpc.CallOption) (*Empty, error)
		DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultUser struct {
		cli zrpc.Client
	}
)

func NewUser(cli zrpc.Client) User {
	return &defaultUser{
		cli: cli,
	}
}

func (m *defaultUser) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUser) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUser) GetAddresses(ctx context.Context, in *GetAddressesRequest, opts ...grpc.CallOption) (*GetAddressesResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetAddresses(ctx, in, opts...)
}

func (m *defaultUser) AddAddress(ctx context.Context, in *AddAddressRequest, opts ...grpc.CallOption) (*AddAddressResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddAddress(ctx, in, opts...)
}

func (m *defaultUser) EditAddress(ctx context.Context, in *EditAddressRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditAddress(ctx, in, opts...)
}

func (m *defaultUser) DeleteAddress(ctx context.Context, in *DeleteAddressRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteAddress(ctx, in, opts...)
}

func (m *defaultUser) SetDefaultAddress(ctx context.Context, in *SetDefaultAddressRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.SetDefaultAddress(ctx, in, opts...)
}

func (m *defaultUser) GetLoginHistory(ctx context.Context, in *GetLoginHistoryRequest, opts ...grpc.CallOption) (*GetLoginHistoryResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetLoginHistory(ctx, in, opts...)
}

func (m *defaultUser) GetUserInfo(ctx context.Context, in *GetUserInfoRequest, opts ...grpc.CallOption) (*GetUserInfoResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}

func (m *defaultUser) GetCollection(ctx context.Context, in *GetCollectionRequest, opts ...grpc.CallOption) (*GetCollectionResponse, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.GetCollection(ctx, in, opts...)
}

func (m *defaultUser) AddCollection(ctx context.Context, in *AddCollectionRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.AddCollection(ctx, in, opts...)
}

func (m *defaultUser) EditCollection(ctx context.Context, in *EditCollectionRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.EditCollection(ctx, in, opts...)
}

func (m *defaultUser) DeleteCollection(ctx context.Context, in *DeleteCollectionRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := user.NewUserClient(m.cli.Conn())
	return client.DeleteCollection(ctx, in, opts...)
}
