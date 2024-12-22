package admin_product

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑商品
func NewEditProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditProductLogic {
	return &EditProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditProductLogic) EditProduct(req *types.EditProductReq) (resp *types.BaseResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
