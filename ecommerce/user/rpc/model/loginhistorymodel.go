package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ LoginHistoryModel = (*customLoginHistoryModel)(nil)

type (
	LoginHistoryModel interface {
		loginHistoryModel
		FindByUserId(ctx context.Context, userId int64, page, pageSize int) ([]*LoginHistory, error)
		CountByUserId(ctx context.Context, userId int64) (int64, error)
		FindByTimeRange(ctx context.Context, userId int64, startTime, endTime time.Time) ([]*LoginHistory, error)
		FindByIp(ctx context.Context, userId int64, ip string) ([]*LoginHistory, error)
		GetLastLogin(ctx context.Context, userId int64) (*LoginHistory, error)
		DeleteByUserId(ctx context.Context, userId int64) error
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

// FindByUserId 分页获取用户的登录历史
func (m *customLoginHistoryModel) FindByUserId(ctx context.Context, userId int64, page, pageSize int) ([]*LoginHistory, error) {
	var histories []*LoginHistory
	offset := (page - 1) * pageSize
	query := fmt.Sprintf("select %s from %s where user_id = ? order by login_time desc limit ? offset ?",
		loginHistoryRows, m.table)

	err := m.QueryRowsNoCacheCtx(ctx, &histories, query, userId, pageSize, offset)
	switch err {
	case nil:
		return histories, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// CountByUserId 获取用户的登录历史总数
func (m *customLoginHistoryModel) CountByUserId(ctx context.Context, userId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where user_id = ?", m.table)

	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// FindByTimeRange 获取指定时间范围内的登录历史
func (m *customLoginHistoryModel) FindByTimeRange(ctx context.Context, userId int64, startTime, endTime time.Time) ([]*LoginHistory, error) {
	var histories []*LoginHistory
	query := fmt.Sprintf("select %s from %s where user_id = ? and login_time between ? and ? order by login_time desc",
		loginHistoryRows, m.table)

	err := m.QueryRowsNoCacheCtx(ctx, &histories, query, userId, startTime, endTime)
	switch err {
	case nil:
		return histories, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// FindByIp 获取指定IP的登录历史
func (m *customLoginHistoryModel) FindByIp(ctx context.Context, userId int64, ip string) ([]*LoginHistory, error) {
	var histories []*LoginHistory
	query := fmt.Sprintf("select %s from %s where user_id = ? and login_ip = ? order by login_time desc",
		loginHistoryRows, m.table)

	err := m.QueryRowsNoCacheCtx(ctx, &histories, query, userId, ip)
	switch err {
	case nil:
		return histories, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// GetLastLogin 获取用户最后一次登录记录
func (m *customLoginHistoryModel) GetLastLogin(ctx context.Context, userId int64) (*LoginHistory, error) {
	var history LoginHistory
	query := fmt.Sprintf("select %s from %s where user_id = ? order by login_time desc limit 1",
		loginHistoryRows, m.table)

	err := m.QueryRowNoCacheCtx(ctx, &history, query, userId)
	switch err {
	case nil:
		return &history, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// DeleteByUserId 删除用户的所有登录历史
func (m *customLoginHistoryModel) DeleteByUserId(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, userId)
	})
	return err
}
