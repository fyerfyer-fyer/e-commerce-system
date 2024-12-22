package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql struct {
		DataSource string
	}

	CacheRedis cache.CacheConf
	BizRedis   cache.CacheConf
	ProductRpc zrpc.RpcClientConf
	Salt       string
	Auth       struct {
		AccessSecret string
		AccessExpire int64
	}
	Limit struct {
		Rate  int
		Burst int
	}
}
