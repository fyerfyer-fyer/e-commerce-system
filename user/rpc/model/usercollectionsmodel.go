package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserCollectionsModel = (*customUserCollectionsModel)(nil)

type (
	// UserCollectionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCollectionsModel.
	UserCollectionsModel interface {
		userCollectionsModel
		withSession(session sqlx.Session) UserCollectionsModel
	}

	customUserCollectionsModel struct {
		*defaultUserCollectionsModel
	}
)

// NewUserCollectionsModel returns a model for the database table.
func NewUserCollectionsModel(conn sqlx.SqlConn) UserCollectionsModel {
	return &customUserCollectionsModel{
		defaultUserCollectionsModel: newUserCollectionsModel(conn),
	}
}

func (m *customUserCollectionsModel) withSession(session sqlx.Session) UserCollectionsModel {
	return NewUserCollectionsModel(sqlx.NewSqlConnFromSession(session))
}
