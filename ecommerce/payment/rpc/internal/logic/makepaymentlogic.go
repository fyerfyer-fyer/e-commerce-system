package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/payment/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/payment/rpc/payment"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakePaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakePaymentLogic {
	return &MakePaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakePaymentLogic) MakePayment(in *payment.PaymentRequest) (*payment.PaymentResponse, error) {
	// todo: add your logic here and delete this line

	return &payment.PaymentResponse{}, nil
}
