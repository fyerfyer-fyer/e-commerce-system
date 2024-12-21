package logic

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductsLogic {
	return &SearchProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SearchProductsLogic) SearchProducts(in *product.SearchProductsRequest) (*product.GetProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.GetProductsResponse{}, nil
}
