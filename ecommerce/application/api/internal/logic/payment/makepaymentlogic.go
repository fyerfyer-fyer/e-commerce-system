package payment

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MakePaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发起支付
func NewMakePaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakePaymentLogic {
	return &MakePaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MakePaymentLogic) MakePayment(req *types.MakePaymentReq) (resp *types.PaymentResp, err error) {
	// todo: add your logic here and delete this line

	return
}
