package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ AdminCodesModel = (*customAdminCodesModel)(nil)

type (
	// AdminCodesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customAdminCodesModel.
	AdminCodesModel interface {
		adminCodesModel
		withSession(session sqlx.Session) AdminCodesModel
	}

	customAdminCodesModel struct {
		*defaultAdminCodesModel
	}
)

// NewAdminCodesModel returns a model for the database table.
func NewAdminCodesModel(conn sqlx.SqlConn) AdminCodesModel {
	return &customAdminCodesModel{
		defaultAdminCodesModel: newAdminCodesModel(conn),
	}
}

func (m *customAdminCodesModel) withSession(session sqlx.Session) AdminCodesModel {
	return NewAdminCodesModel(sqlx.NewSqlConnFromSession(session))
}
