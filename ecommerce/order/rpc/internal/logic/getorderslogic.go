package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrdersLogic {
	return &GetOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrdersLogic) GetOrders(in *order.GetOrdersRequest) (*order.GetOrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrdersResponse{}, nil
}
