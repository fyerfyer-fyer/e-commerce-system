// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: product.proto

package server

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/logic"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"
)

type ProductsServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductsServer
}

func NewProductsServer(svcCtx *svc.ServiceContext) *ProductsServer {
	return &ProductsServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductsServer) GetProduct(ctx context.Context, in *product.GetProductRequest) (*product.Product, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

func (s *ProductsServer) GetProductsByCategory(ctx context.Context, in *product.GetProductsByCategoryRequest) (*product.GetProductsResponse, error) {
	l := logic.NewGetProductsByCategoryLogic(ctx, s.svcCtx)
	return l.GetProductsByCategory(in)
}

func (s *ProductsServer) SearchProducts(ctx context.Context, in *product.SearchProductsRequest) (*product.GetProductsResponse, error) {
	l := logic.NewSearchProductsLogic(ctx, s.svcCtx)
	return l.SearchProducts(in)
}

func (s *ProductsServer) UpdateProductStock(ctx context.Context, in *product.UpdateProductStockRequest) (*product.Empty, error) {
	l := logic.NewUpdateProductStockLogic(ctx, s.svcCtx)
	return l.UpdateProductStock(in)
}
