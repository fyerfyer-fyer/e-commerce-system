// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.3

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	adminCodesFieldNames          = builder.RawFieldNames(&AdminCodes{})
	adminCodesRows                = strings.Join(adminCodesFieldNames, ",")
	adminCodesRowsExpectAutoSet   = strings.Join(stringx.Remove(adminCodesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	adminCodesRowsWithPlaceHolder = strings.Join(stringx.Remove(adminCodesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheUserServiceAdminCodesIdPrefix   = "cache:userService:adminCodes:id:"
	cacheUserServiceAdminCodesCodePrefix = "cache:userService:adminCodes:code:"
)

type (
	adminCodesModel interface {
		Insert(ctx context.Context, data *AdminCodes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*AdminCodes, error)
		FindOneByCode(ctx context.Context, code string) (*AdminCodes, error)
		Update(ctx context.Context, data *AdminCodes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAdminCodesModel struct {
		sqlc.CachedConn
		table string
	}

	AdminCodes struct {
		Id        int64        `db:"id"`         // Admin Code ID
		Code      string       `db:"code"`       // The actual admin code
		CreatedAt time.Time    `db:"created_at"` // Creation time
		Used      int64        `db:"used"`       // is code used
		ExpiredAt sql.NullTime `db:"expired_at"` // code expired time
	}
)

func newAdminCodesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultAdminCodesModel {
	return &defaultAdminCodesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`admin_codes`",
	}
}

func (m *defaultAdminCodesModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	userServiceAdminCodesCodeKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesCodePrefix, data.Code)
	userServiceAdminCodesIdKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, userServiceAdminCodesCodeKey, userServiceAdminCodesIdKey)
	return err
}

func (m *defaultAdminCodesModel) FindOne(ctx context.Context, id int64) (*AdminCodes, error) {
	userServiceAdminCodesIdKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesIdPrefix, id)
	var resp AdminCodes
	err := m.QueryRowCtx(ctx, &resp, userServiceAdminCodesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminCodesRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminCodesModel) FindOneByCode(ctx context.Context, code string) (*AdminCodes, error) {
	userServiceAdminCodesCodeKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesCodePrefix, code)
	var resp AdminCodes
	err := m.QueryRowIndexCtx(ctx, &resp, userServiceAdminCodesCodeKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `code` = ? limit 1", adminCodesRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, code); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAdminCodesModel) Insert(ctx context.Context, data *AdminCodes) (sql.Result, error) {
	userServiceAdminCodesCodeKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesCodePrefix, data.Code)
	userServiceAdminCodesIdKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, adminCodesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Code, data.Used, data.ExpiredAt)
	}, userServiceAdminCodesCodeKey, userServiceAdminCodesIdKey)
	return ret, err
}

func (m *defaultAdminCodesModel) Update(ctx context.Context, newData *AdminCodes) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	userServiceAdminCodesCodeKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesCodePrefix, data.Code)
	userServiceAdminCodesIdKey := fmt.Sprintf("%s%v", cacheUserServiceAdminCodesIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, adminCodesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Code, newData.Used, newData.ExpiredAt, newData.Id)
	}, userServiceAdminCodesCodeKey, userServiceAdminCodesIdKey)
	return err
}

func (m *defaultAdminCodesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheUserServiceAdminCodesIdPrefix, primary)
}

func (m *defaultAdminCodesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", adminCodesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultAdminCodesModel) tableName() string {
	return m.table
}