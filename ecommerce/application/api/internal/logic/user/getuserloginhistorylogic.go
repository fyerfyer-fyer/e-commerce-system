package user

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLoginHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户登录历史
func NewGetUserLoginHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginHistoryLogic {
	return &GetUserLoginHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLoginHistoryLogic) GetUserLoginHistory(req *types.LoginHistoryReq) (resp *types.LoginHistoryResp, err error) {
	// todo: add your logic here and delete this line

	return
}
