package user

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加用户地址
func NewAddUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserAddressLogic {
	return &AddUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserAddressLogic) AddUserAddress(req *types.AddAddressReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
