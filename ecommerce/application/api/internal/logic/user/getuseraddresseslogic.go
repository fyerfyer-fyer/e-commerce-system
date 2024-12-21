package user

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户地址列表
func NewGetUserAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressesLogic {
	return &GetUserAddressesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAddressesLogic) GetUserAddresses(req *types.Empty) (resp *types.AddressListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
