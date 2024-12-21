package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditAddressLogic {
	return &EditAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditAddressLogic) EditAddress(in *user.EditAddressRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
