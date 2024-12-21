package seckill

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeckillEventsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取秒杀活动列表
func NewGetSeckillEventsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillEventsLogic {
	return &GetSeckillEventsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSeckillEventsLogic) GetSeckillEvents(req *types.GetSeckillEventsReq) (resp *types.GetSeckillEventsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
