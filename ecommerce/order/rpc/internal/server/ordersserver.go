// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: order.proto

package server

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/logic"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/internal/svc"
	"github.com/fyerfyer/e-commerce-system/ecommerce/order/rpc/order"
)

type OrdersServer struct {
	svcCtx *svc.ServiceContext
	order.UnimplementedOrdersServer
}

func NewOrdersServer(svcCtx *svc.ServiceContext) *OrdersServer {
	return &OrdersServer{
		svcCtx: svcCtx,
	}
}

func (s *OrdersServer) GetOrders(ctx context.Context, in *order.GetOrdersRequest) (*order.GetOrdersResponse, error) {
	l := logic.NewGetOrdersLogic(ctx, s.svcCtx)
	return l.GetOrders(in)
}

func (s *OrdersServer) GetOrder(ctx context.Context, in *order.GetOrderRequest) (*order.Order, error) {
	l := logic.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}

func (s *OrdersServer) SubmitOrder(ctx context.Context, in *order.SubmitOrderRequest) (*order.SubmitOrderResponse, error) {
	l := logic.NewSubmitOrderLogic(ctx, s.svcCtx)
	return l.SubmitOrder(in)
}

func (s *OrdersServer) EditOrder(ctx context.Context, in *order.EditOrderRequest) (*order.Empty, error) {
	l := logic.NewEditOrderLogic(ctx, s.svcCtx)
	return l.EditOrder(in)
}

func (s *OrdersServer) DeleteOrder(ctx context.Context, in *order.DeleteOrderRequest) (*order.Empty, error) {
	l := logic.NewDeleteOrderLogic(ctx, s.svcCtx)
	return l.DeleteOrder(in)
}

func (s *OrdersServer) GetOrderItems(ctx context.Context, in *order.GetOrderItemsRequest) (*order.GetOrderItemsResponse, error) {
	l := logic.NewGetOrderItemsLogic(ctx, s.svcCtx)
	return l.GetOrderItems(in)
}