package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressesLogic {
	return &GetAddressesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAddressesLogic) GetAddresses(in *user.GetAddressesRequest) (*user.GetAddressesResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetAddressesResponse{}, nil
}
