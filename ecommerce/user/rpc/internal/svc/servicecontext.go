package svc

import (
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/internal/config"
	"github.com/fyerfyer/e-commerce-system/ecommerce/user/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	// MySQL models
	UsersModel           model.UsersModel
	UserAddressesModel   model.UserAddressesModel
	AdminCodesModel      model.AdminCodesModel
	LoginHistoryModel    model.LoginHistoryModel
	UserCollectionsModel model.UserCollectionsModel

	// Redis clients
	BizRedis *redis.Redis // 业务缓存

	// RPC clients
	ProductRpc zrpc.Client // 商品服务客户端
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化MySQL连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	// 初始化Redis
	bizRedis := redis.New(c.BizRedis[0].Host, func(r *redis.Redis) {
		r.Type = c.BizRedis[0].Type
		r.Pass = c.BizRedis[0].Pass
	})

	return &ServiceContext{
		Config: c,

		// 初始化所有Model，使用Redis缓存
		UsersModel:           model.NewUsersModel(conn, c.CacheRedis),
		UserAddressesModel:   model.NewUserAddressesModel(conn, c.CacheRedis),
		AdminCodesModel:      model.NewAdminCodesModel(conn, c.CacheRedis),
		LoginHistoryModel:    model.NewLoginHistoryModel(conn, c.CacheRedis),
		UserCollectionsModel: model.NewUserCollectionsModel(conn, c.CacheRedis),

		// Redis客户端
		BizRedis: bizRedis,

		// RPC客户端
		ProductRpc: zrpc.MustNewClient(c.ProductRpc),
	}
}
