package seckill

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提交秒杀订单
func NewPlaceSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceSeckillOrderLogic {
	return &PlaceSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceSeckillOrderLogic) PlaceSeckillOrder(req *types.SeckillOrderReq) (resp *types.SeckillOrderResp, err error) {
	// todo: add your logic here and delete this line

	return
}
