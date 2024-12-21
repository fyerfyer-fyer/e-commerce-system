package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductsByCategoryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductsByCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductsByCategoryLogic {
	return &GetProductsByCategoryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductsByCategoryLogic) GetProductsByCategory(in *product.GetProductsByCategoryRequest) (*product.GetProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductsResponse{}, nil
}
