package order

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取订单详情
func NewGetOrderDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailsLogic {
	return &GetOrderDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderDetailsLogic) GetOrderDetails(req *types.Empty) (resp *types.OrderDetailsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
