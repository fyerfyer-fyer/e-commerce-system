package order

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserOrdersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户订单列表
func NewGetUserOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserOrdersLogic {
	return &GetUserOrdersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserOrdersLogic) GetUserOrders(req *types.PaginationReq) (resp *types.OrderListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
