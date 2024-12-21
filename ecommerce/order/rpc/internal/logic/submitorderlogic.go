package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitOrderLogic {
	return &SubmitOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitOrderLogic) SubmitOrder(in *order.SubmitOrderRequest) (*order.SubmitOrderResponse, error) {
	// todo: add your logic here and delete this line

	return &order.SubmitOrderResponse{}, nil
}
