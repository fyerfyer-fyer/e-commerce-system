package cart

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取购物车列表
func NewGetCartItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartItemsLogic {
	return &GetCartItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCartItemsLogic) GetCartItems(req *types.Empty) (resp *types.CartListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
