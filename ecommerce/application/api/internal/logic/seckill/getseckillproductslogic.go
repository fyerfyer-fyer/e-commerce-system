package seckill

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSeckillProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取秒杀商品列表
func NewGetSeckillProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSeckillProductsLogic {
	return &GetSeckillProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSeckillProductsLogic) GetSeckillProducts(req *types.GetSeckillProductsReq) (resp *types.GetSeckillProductsResp, err error) {
	// todo: add your logic here and delete this line

	return
}
