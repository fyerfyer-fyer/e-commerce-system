package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderItemsLogic {
	return &GetOrderItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderItemsLogic) GetOrderItems(in *order.GetOrderItemsRequest) (*order.GetOrderItemsResponse, error) {
	// todo: add your logic here and delete this line

	return &order.GetOrderItemsResponse{}, nil
}
