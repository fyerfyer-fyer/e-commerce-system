package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/payment/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/payment/rpc/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundLogic {
	return &RefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefundLogic) Refund(in *payment.RefundRequest) (*payment.RefundResponse, error) {
	// todo: add your logic here and delete this line

	return &payment.RefundResponse{}, nil
}
