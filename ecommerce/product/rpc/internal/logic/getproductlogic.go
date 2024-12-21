package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductRequest) (*product.Product, error) {
	// todo: add your logic here and delete this line

	return &product.Product{}, nil
}
