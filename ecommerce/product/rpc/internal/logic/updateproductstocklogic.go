package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductStockLogic {
	return &UpdateProductStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductStockLogic) UpdateProductStock(in *product.UpdateProductStockRequest) (*product.Empty, error) {
	// todo: add your logic here and delete this line

	return &product.Empty{}, nil
}
