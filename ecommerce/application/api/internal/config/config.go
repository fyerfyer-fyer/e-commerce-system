package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf

	// JWT认证配置
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 管理员JWT认证配置
	AdminAuth struct {
		AccessSecret string
		AccessExpire int64
	}

	// 无需认证的配置
	None struct {
		AccessSecret string
		AccessExpire int64
	}

	// 各个RPC服务配置
	UserRpc    zrpc.RpcClientConf
	OrderRpc   zrpc.RpcClientConf
	ProductRpc zrpc.RpcClientConf
	CartRpc    zrpc.RpcClientConf
	PaymentRpc zrpc.RpcClientConf
	CommentRpc zrpc.RpcClientConf
	SeckillRpc zrpc.RpcClientConf

	// Redis配置
	Redis redis.RedisConf

	// 中间件配置
	Middleware struct {
		Limit struct {
			Rate              int64 // 每秒请求数
			Burst             int64 // 突发请求数
			EnableLocalLimit  bool  // 启用本地限流
			EnableGlobalLimit bool  // 启用全局限流
		}
	}

	// 跨域配置
	CorsConfig struct {
		Enable           bool
		AllowOrigins     string
		AllowMethods     string
		AllowHeaders     string
		ExposeHeaders    string
		AllowCredentials bool
	}

	// 链路追踪配置
	Telemetry struct {
		Name     string
		Endpoint string
		Sampler  float64
		Batcher  string
	}

	// 超时配置
	Timeout struct {
		Rpc  int64 // RPC调用超时时间（毫秒）
		Http int64 // HTTP请求超时时间（毫秒）
	}
}
