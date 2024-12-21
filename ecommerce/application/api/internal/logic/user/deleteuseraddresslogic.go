package user

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除用户地址
func NewDeleteUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAddressLogic {
	return &DeleteUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteUserAddressLogic) DeleteUserAddress(req *types.Empty) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
