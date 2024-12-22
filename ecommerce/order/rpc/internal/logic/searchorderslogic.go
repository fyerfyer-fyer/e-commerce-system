package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchOrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchOrdersLogic {
	return &SearchOrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchOrdersLogic) SearchOrders(in *order.SearchOrdersRequest) (*order.SearchOrdersResponse, error) {
	// todo: add your logic here and delete this line

	return &order.SearchOrdersResponse{}, nil
}
