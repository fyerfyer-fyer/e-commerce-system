package user

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditUserAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑用户地址
func NewEditUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditUserAddressLogic {
	return &EditUserAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditUserAddressLogic) EditUserAddress(req *types.EditAddressReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
