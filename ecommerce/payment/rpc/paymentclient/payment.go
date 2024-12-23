// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: payment.proto

package paymentclient

import (
	"context"

	"github.com/fyerfyer/e-commerce-system/ecommerce/payment/rpc/payment"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Empty                    = payment.Empty
	GetPaymentStatusRequest  = payment.GetPaymentStatusRequest
	GetPaymentStatusResponse = payment.GetPaymentStatusResponse
	PaymentRequest           = payment.PaymentRequest
	PaymentResponse          = payment.PaymentResponse
	RefundRequest            = payment.RefundRequest
	RefundResponse           = payment.RefundResponse

	Payment interface {
		MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error)
		GetPaymentStatus(ctx context.Context, in *GetPaymentStatusRequest, opts ...grpc.CallOption) (*GetPaymentStatusResponse, error)
		Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error)
	}

	defaultPayment struct {
		cli zrpc.Client
	}
)

func NewPayment(cli zrpc.Client) Payment {
	return &defaultPayment{
		cli: cli,
	}
}

func (m *defaultPayment) MakePayment(ctx context.Context, in *PaymentRequest, opts ...grpc.CallOption) (*PaymentResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.MakePayment(ctx, in, opts...)
}

func (m *defaultPayment) GetPaymentStatus(ctx context.Context, in *GetPaymentStatusRequest, opts ...grpc.CallOption) (*GetPaymentStatusResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.GetPaymentStatus(ctx, in, opts...)
}

func (m *defaultPayment) Refund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error) {
	client := payment.NewPaymentClient(m.cli.Conn())
	return client.Refund(ctx, in, opts...)
}
