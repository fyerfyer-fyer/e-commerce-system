package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminCodesModel = (*customAdminCodesModel)(nil)

type (
	// AdminCodesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminCodesModel.
	AdminCodesModel interface {
		adminCodesModel
	}

	customAdminCodesModel struct {
		*defaultAdminCodesModel
	}
)

// NewAdminCodesModel returns a model for the database table.
func NewAdminCodesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) AdminCodesModel {
	return &customAdminCodesModel{
		defaultAdminCodesModel: newAdminCodesModel(conn, c, opts...),
	}
}
