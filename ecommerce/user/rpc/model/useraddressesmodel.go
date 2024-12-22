package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAddressesModel = (*customUserAddressesModel)(nil)

type (
	UserAddressesModel interface {
		userAddressesModel
		FindByUserId(ctx context.Context, userId int64) ([]*UserAddresses, error)
		SetDefaultAddress(ctx context.Context, userId, addressId int64) error
		GetDefaultAddress(ctx context.Context, userId int64) (*UserAddresses, error)
		VerifyAddressOwnership(ctx context.Context, userId, addressId int64) (bool, error)
		CountUserAddresses(ctx context.Context, userId int64) (int64, error)
		DeleteUserAddresses(ctx context.Context, userId int64) error
	}

	customUserAddressesModel struct {
		*defaultUserAddressesModel
	}
)

// NewUserAddressesModel returns a model for the database table.
func NewUserAddressesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAddressesModel {
	return &customUserAddressesModel{
		defaultUserAddressesModel: newUserAddressesModel(conn, c, opts...),
	}
}

// FindByUserId 获取用户的所有地址
func (m *customUserAddressesModel) FindByUserId(ctx context.Context, userId int64) ([]*UserAddresses, error) {
	var addresses []*UserAddresses
	query := fmt.Sprintf("select %s from %s where user_id = ? order by is_default desc, id desc", userAddressesRows, m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &addresses, query, userId)
	switch err {
	case nil:
		return addresses, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// SetDefaultAddress 设置默认地址
func (m *customUserAddressesModel) SetDefaultAddress(ctx context.Context, userId, addressId int64) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		// 先将该用户的所有地址设置为非默认
		_, err := session.ExecCtx(ctx,
			fmt.Sprintf("update %s set is_default = 0 where user_id = ?", m.table),
			userId)
		if err != nil {
			return err
		}

		// 将指定地址设置为默认
		_, err = session.ExecCtx(ctx,
			fmt.Sprintf("update %s set is_default = 1 where id = ? and user_id = ?", m.table),
			addressId, userId)
		return err
	})
}

// GetDefaultAddress 获取用户的默认地址
func (m *customUserAddressesModel) GetDefaultAddress(ctx context.Context, userId int64) (*UserAddresses, error) {
	var address UserAddresses
	query := fmt.Sprintf("select %s from %s where user_id = ? and is_default = 1 limit 1", userAddressesRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &address, query, userId)
	switch err {
	case nil:
		return &address, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

// VerifyAddressOwnership 验证地址是否属于指定用户
func (m *customUserAddressesModel) VerifyAddressOwnership(ctx context.Context, userId, addressId int64) (bool, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where id = ? and user_id = ?", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, addressId, userId)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// CountUserAddresses 获取用户地址数量
func (m *customUserAddressesModel) CountUserAddresses(ctx context.Context, userId int64) (int64, error) {
	var count int64
	query := fmt.Sprintf("select count(*) from %s where user_id = ?", m.table)
	err := m.QueryRowNoCacheCtx(ctx, &count, query, userId)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteUserAddresses 删除用户的所有地址
func (m *customUserAddressesModel) DeleteUserAddresses(ctx context.Context, userId int64) error {
	query := fmt.Sprintf("delete from %s where user_id = ?", m.table)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, userId)
	})
	return err
}
