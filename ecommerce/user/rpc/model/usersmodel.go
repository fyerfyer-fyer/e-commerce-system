package model

import (
    "context"
    "database/sql"
    "fmt"
    "github.com/zeromicro/go-zero/core/stores/cache"
    "github.com/zeromicro/go-zero/core/stores/sqlx"
    "strings"
)

var _ UsersModel = (*customUsersModel)(nil)

type (
    UsersModel interface {
        usersModel
        FindByIdentifier(ctx context.Context, identifier string) (*Users, error)
        Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
        FindByIds(ctx context.Context, ids []int64) ([]*Users, error)
        UpdateRole(ctx context.Context, id int64, role string) error
        VerifyPassword(ctx context.Context, id int64, password string) (bool, error)
    }

    customUsersModel struct {
        *defaultUsersModel
    }
)

func NewUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UsersModel {
    return &customUsersModel{
        defaultUsersModel: newUsersModel(conn, c, opts...),
    }
}

func (m *customUsersModel) FindByIdentifier(ctx context.Context, identifier string) (*Users, error) {
    var user Users
    query := fmt.Sprintf("select %s from %s where username = ? or email = ? or phone = ? limit 1", usersRows, m.table)
    err := m.QueryRowNoCacheCtx(ctx, &user, query, identifier, identifier, identifier)
    switch err {
    case nil:
        return &user, nil
    case sqlx.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

func (m *customUsersModel) Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
    return m.TransactCtx(ctx, fn)
}

func (m *customUsersModel) FindByIds(ctx context.Context, ids []int64) ([]*Users, error) {
    var users []*Users
    // 构建 IN 查询的占位符
    placeholders := make([]string, len(ids))
    args := make([]interface{}, len(ids))
    for i, id := range ids {
        placeholders[i] = "?"
        args[i] = id
    }
    
    query := fmt.Sprintf("select %s from %s where id in (%s)", 
        usersRows, m.table, strings.Join(placeholders, ","))
    
    err := m.QueryRowsNoCacheCtx(ctx, &users, query, args...)
    switch err {
    case nil:
        return users, nil
    case sqlx.ErrNotFound:
        return nil, ErrNotFound
    default:
        return nil, err
    }
}

func (m *customUsersModel) UpdateRole(ctx context.Context, id int64, role string) error {
    userIdKey := fmt.Sprintf("%s%v", cacheUserServiceUsersIdPrefix, id)
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
        query := fmt.Sprintf("update %s set role = ? where id = ?", m.table)
        return conn.ExecCtx(ctx, query, role, id)
    }, userIdKey)
    return err
}

func (m *customUsersModel) VerifyPassword(ctx context.Context, id int64, password string) (bool, error) {
    var count int
    query := fmt.Sprintf("select count(*) from %s where id = ? and password = ?", m.table)
    err := m.QueryRowNoCacheCtx(ctx, &count, query, id, password)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}