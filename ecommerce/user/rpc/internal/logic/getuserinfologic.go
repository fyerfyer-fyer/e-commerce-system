package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetUserInfoResponse{}, nil
}
