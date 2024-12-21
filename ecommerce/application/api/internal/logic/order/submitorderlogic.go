package order

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提交订单
func NewSubmitOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitOrderLogic {
	return &SubmitOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitOrderLogic) SubmitOrder(req *types.SubmitOrderReq) (resp *types.SubmitOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
