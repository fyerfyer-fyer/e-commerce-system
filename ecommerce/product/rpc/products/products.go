// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: product.proto

package products

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/product/rpc/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Category                     = product.Category
	Empty                        = product.Empty
	GetProductRequest            = product.GetProductRequest
	GetProductsByCategoryRequest = product.GetProductsByCategoryRequest
	GetProductsResponse          = product.GetProductsResponse
	Product                      = product.Product
	ProductAttribute             = product.ProductAttribute
	SearchProductsRequest        = product.SearchProductsRequest
	UpdateProductStockRequest    = product.UpdateProductStockRequest

	Products interface {
		GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error)
		GetProductsByCategory(ctx context.Context, in *GetProductsByCategoryRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
		SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error)
		UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*Empty, error)
	}

	defaultProducts struct {
		cli zrpc.Client
	}
)

func NewProducts(cli zrpc.Client) Products {
	return &defaultProducts{
		cli: cli,
	}
}

func (m *defaultProducts) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*Product, error) {
	client := product.NewProductsClient(m.cli.Conn())
	return client.GetProduct(ctx, in, opts...)
}

func (m *defaultProducts) GetProductsByCategory(ctx context.Context, in *GetProductsByCategoryRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	client := product.NewProductsClient(m.cli.Conn())
	return client.GetProductsByCategory(ctx, in, opts...)
}

func (m *defaultProducts) SearchProducts(ctx context.Context, in *SearchProductsRequest, opts ...grpc.CallOption) (*GetProductsResponse, error) {
	client := product.NewProductsClient(m.cli.Conn())
	return client.SearchProducts(ctx, in, opts...)
}

func (m *defaultProducts) UpdateProductStock(ctx context.Context, in *UpdateProductStockRequest, opts ...grpc.CallOption) (*Empty, error) {
	client := product.NewProductsClient(m.cli.Conn())
	return client.UpdateProductStock(ctx, in, opts...)
}
