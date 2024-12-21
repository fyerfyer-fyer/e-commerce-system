package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectionLogic {
	return &GetCollectionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCollectionLogic) GetCollection(in *user.GetCollectionRequest) (*user.GetCollectionResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetCollectionResponse{}, nil
}
