package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLoginHistoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLoginHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLoginHistoryLogic {
	return &GetLoginHistoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLoginHistoryLogic) GetLoginHistory(in *user.GetLoginHistoryRequest) (*user.GetLoginHistoryResponse, error) {
	// todo: add your logic here and delete this line

	return &user.GetLoginHistoryResponse{}, nil
}
