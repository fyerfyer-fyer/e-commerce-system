package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserCollectionsModel = (*customUserCollectionsModel)(nil)

type (
	// UserCollectionsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserCollectionsModel.
	UserCollectionsModel interface {
		userCollectionsModel
		GetUserCollections(ctx context.Context, userId int64) ([]int64, error)
		ClearUserCollections(ctx context.Context, userId int64) error
		BatchInsert(ctx context.Context, userId int64, productIds []int64) error
		DeleteByUserIdAndProductId(ctx context.Context, userId, productId int64) error
		CountUserCollections(ctx context.Context, userId int64) (int64, error)
		GetUserCollectionsWithPage(ctx context.Context, userId int64, page, pageSize int) ([]int64, error)
		IsCollected(ctx context.Context, userId, productId int64) (bool, error)
	}

	customUserCollectionsModel struct {
		*defaultUserCollectionsModel
	}
)

// NewUserCollectionsModel returns a model for the database table.
func NewUserCollectionsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserCollectionsModel {
	return &customUserCollectionsModel{
		defaultUserCollectionsModel: newUserCollectionsModel(conn, c, opts...),
	}
}

func (m *customUserCollectionsModel) GetUserCollections(ctx context.Context, userId int64) ([]int64, error) {
	var productIds []int64
	query := `SELECT product_id FROM user_collections WHERE user_id = ?`
	err := m.QueryRowsNoCacheCtx(ctx, &productIds, query, userId)
	switch err {
	case nil:
		return productIds, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *customUserCollectionsModel) ClearUserCollections(ctx context.Context, userId int64) error {
	clearQuery := `DELETE FROM user_collections WHERE user_id = ?`
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		return conn.ExecCtx(ctx, clearQuery, userId)
	})
	return err
}

// BatchInsert 批量插入收藏
func (m *customUserCollectionsModel) BatchInsert(ctx context.Context, userId int64, productIds []int64) error {
    if len(productIds) == 0 {
        return nil
    }
    
    var values []string
    var args []interface{}
    
    for _, productId := range productIds {
        values = append(values, "(?, ?)")
        args = append(args, userId, productId)
    }
    
    query := fmt.Sprintf("INSERT IGNORE INTO user_collections (user_id, product_id) VALUES %s",
        strings.Join(values, ","))
    
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return conn.ExecCtx(ctx, query, args...)
    })
    return err
}

// DeleteByUserIdAndProductId 删除指定用户的指定商品收藏
func (m *customUserCollectionsModel) DeleteByUserIdAndProductId(ctx context.Context, userId, productId int64) error {
    query := `DELETE FROM user_collections WHERE user_id = ? AND product_id = ?`
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        return conn.ExecCtx(ctx, query, userId, productId)
    })
    return err
}

// CountUserCollections 获取用户收藏总数
func (m *customUserCollectionsModel) CountUserCollections(ctx context.Context, userId int64) (int64, error) {
    var count int64
    query := `SELECT COUNT(*) FROM user_collections WHERE user_id = ?`
    err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
    if err != nil {
        return 0, err
    }
    return count, nil
}

// GetUserCollectionsWithPage 分页获取用户收藏
func (m *customUserCollectionsModel) GetUserCollectionsWithPage(ctx context.Context, userId int64, page, pageSize int) ([]int64, error) {
    var productIds []int64
    offset := (page - 1) * pageSize
    query := `SELECT product_id FROM user_collections WHERE user_id = ? ORDER BY created_at DESC LIMIT ? OFFSET ?`
    err := m.QueryRowsNoCacheCtx(ctx, &productIds, query, userId, pageSize, offset)
    switch err {
    case nil:
        return productIds, nil
    case sqlx.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

// IsCollected 检查商品是否已被用户收藏
func (m *customUserCollectionsModel) IsCollected(ctx context.Context, userId, productId int64) (bool, error) {
    var count int64
    query := `SELECT COUNT(*) FROM user_collections WHERE user_id = ? AND product_id = ?`
    err := m.QueryRowNoCacheCtx(ctx, &count, query, userId, productId)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}