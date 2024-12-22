package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ LoginHistoryModel = (*customLoginHistoryModel)(nil)

type (
	// LoginHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginHistoryModel.
	LoginHistoryModel interface {
		loginHistoryModel
	}

	customLoginHistoryModel struct {
		*defaultLoginHistoryModel
	}
)

// NewLoginHistoryModel returns a model for the database table.
func NewLoginHistoryModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) LoginHistoryModel {
	return &customLoginHistoryModel{
		defaultLoginHistoryModel: newLoginHistoryModel(conn, c, opts...),
	}
}
