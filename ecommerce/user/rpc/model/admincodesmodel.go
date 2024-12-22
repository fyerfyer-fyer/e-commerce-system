package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ AdminCodesModel = (*customAdminCodesModel)(nil)

type (
	AdminCodesModel interface {
		adminCodesModel
		FindValidCode(ctx context.Context, code string) (*AdminCodes, error)
		MarkCodeAsUsed(ctx context.Context, code string) error
		BatchGenerateCodes(ctx context.Context, count int, expireTime time.Time) error
		DeleteExpiredCodes(ctx context.Context) error
		GetValidCodes(ctx context.Context, page, pageSize int) ([]*AdminCodes, error)
		CountValidCodes(ctx context.Context) (int64, error)
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

// FindValidCode 查找有效的管理员注册码
func (m *customAdminCodesModel) FindValidCode(ctx context.Context, code string) (*AdminCodes, error) {
	var adminCode AdminCodes
	query := fmt.Sprintf("select %s from %s where code = ? and used = 0 and (expired_at is null or expired_at > ?) limit 1",
		adminCodesRows, m.table)

	err := m.QueryRowNoCacheCtx(ctx, &adminCode, query, code, time.Now())
	switch err {
	case nil:
		return &adminCode, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// MarkCodeAsUsed 将注册码标记为已使用
func (m *customAdminCodesModel) MarkCodeAsUsed(ctx context.Context, code string) error {
	query := fmt.Sprintf("update %s set used = 1 where code = ? and used = 0", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, code)
	})
	return err
}

// BatchGenerateCodes 批量生成管理员注册码
func (m *customAdminCodesModel) BatchGenerateCodes(ctx context.Context, count int, expireTime time.Time) error {
	if count <= 0 {
		return nil
	}

	values := make([]string, count)
	args := make([]interface{}, count*2)

	for i := 0; i < count; i++ {
		values[i] = "(?, ?)"
		// 这里需要生成唯一的注册码，可以使用UUID或其他方式
		code := fmt.Sprintf("ADMIN_%d_%d", time.Now().UnixNano(), i)
		args[i*2] = code
		args[i*2+1] = expireTime
	}

	query := fmt.Sprintf("insert into %s (code, expired_at) values %s",
		m.table, strings.Join(values, ","))

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, args...)
	})
	return err
}

// DeleteExpiredCodes 删除过期的注册码
func (m *customAdminCodesModel) DeleteExpiredCodes(ctx context.Context) error {
	query := fmt.Sprintf("delete from %s where expired_at < ? and used = 0", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, time.Now())
	})
	return err
}

// GetValidCodes 分页获取有效的注册码
func (m *customAdminCodesModel) GetValidCodes(ctx context.Context, page, pageSize int) ([]*AdminCodes, error) {
	var codes []*AdminCodes
	offset := (page - 1) * pageSize

	query := fmt.Sprintf(`
        select %s from %s 
        where used = 0 and (expired_at is null or expired_at > ?) 
        order by id desc limit ? offset ?`,
		adminCodesRows, m.table)

	err := m.QueryRowsNoCacheCtx(ctx, &codes, query, time.Now(), pageSize, offset)
	switch err {
	case nil:
		return codes, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// CountValidCodes 统计有效的注册码数量
func (m *customAdminCodesModel) CountValidCodes(ctx context.Context) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where used = 0 and (expired_at is null or expired_at > ?)",
		m.table)

	err := m.QueryRowNoCacheCtx(ctx, &count, query, time.Now())
	if err != nil {
		return 0, err
	}
	return count, nil
}
