package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ LoginHistoryModel = (*customLoginHistoryModel)(nil)

type (
	// LoginHistoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customLoginHistoryModel.
	LoginHistoryModel interface {
		loginHistoryModel
		withSession(session sqlx.Session) LoginHistoryModel
	}

	customLoginHistoryModel struct {
		*defaultLoginHistoryModel
	}
)

// NewLoginHistoryModel returns a model for the database table.
func NewLoginHistoryModel(conn sqlx.SqlConn) LoginHistoryModel {
	return &customLoginHistoryModel{
		defaultLoginHistoryModel: newLoginHistoryModel(conn),
	}
}

func (m *customLoginHistoryModel) withSession(session sqlx.Session) LoginHistoryModel {
	return NewLoginHistoryModel(sqlx.NewSqlConnFromSession(session))
}
