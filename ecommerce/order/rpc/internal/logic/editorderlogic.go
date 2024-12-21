package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEditOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditOrderLogic {
	return &EditOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EditOrderLogic) EditOrder(in *order.EditOrderRequest) (*order.Empty, error) {
	// todo: add your logic here and delete this line

	return &order.Empty{}, nil
}
