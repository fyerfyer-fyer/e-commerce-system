package svc

import (
	"time"

	"github.com/fyerfyer/e-commerce-system/ecommerce/application/api/internal/config"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	UserRpc    zrpc.Client
	OrderRpc   zrpc.Client
	ProductRpc zrpc.Client
	CartRpc    zrpc.Client
	PaymentRpc zrpc.Client
	CommentRpc zrpc.Client
	SeckillRpc zrpc.Client
	Redis      *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 超时配置
	rpcClientOpts := []zrpc.ClientOption{
		zrpc.WithTimeout(time.Duration(c.Timeout.Rpc) * time.Millisecond),
	}

	return &ServiceContext{
		Config:     c,
		UserRpc:    zrpc.MustNewClient(c.UserRpc, rpcClientOpts...),
		OrderRpc:   zrpc.MustNewClient(c.OrderRpc, rpcClientOpts...),
		ProductRpc: zrpc.MustNewClient(c.ProductRpc, rpcClientOpts...),
		CartRpc:    zrpc.MustNewClient(c.CartRpc, rpcClientOpts...),
		PaymentRpc: zrpc.MustNewClient(c.PaymentRpc, rpcClientOpts...),
		CommentRpc: zrpc.MustNewClient(c.CommentRpc, rpcClientOpts...),
		SeckillRpc: zrpc.MustNewClient(c.SeckillRpc, rpcClientOpts...),
		Redis: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
	}
}
