package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddCollectionLogic {
	return &AddCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddCollectionLogic) AddCollection(in *user.AddCollectionRequest) (*user.Empty, error) {
	// todo: add your logic here and delete this line

	return &user.Empty{}, nil
}
