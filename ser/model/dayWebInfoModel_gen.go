// Code generated by goctl. DO NOT EDIT!

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
	dayWebInfoFieldNames          = builder.RawFieldNames(&DayWebInfo{})
	dayWebInfoRows                = strings.Join(dayWebInfoFieldNames, ",")
	dayWebInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(dayWebInfoFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	dayWebInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(dayWebInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheCloudcompDayWebInfoIdPrefix   = "cache:cloudcomp:dayWebInfo:id:"
	cacheCloudcompDayWebInfoDatePrefix = "cache:cloudcomp:dayWebInfo:date:"
)

type (
	dayWebInfoModel interface {
		Insert(ctx context.Context, data *DayWebInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*DayWebInfo, error)
		FindOneByDate(ctx context.Context, date string) (*DayWebInfo, error)
		Update(ctx context.Context, data *DayWebInfo) error
		Delete(ctx context.Context, id int64) error
	}

	defaultDayWebInfoModel struct {
		sqlc.CachedConn
		table string
	}

	DayWebInfo struct {
		Id         int64     `db:"id"`
		Date       string    `db:"date"`        // 时间
		Login      int64     `db:"login"`       // 登录数量
		Lookup     int64     `db:"lookup"`      // 网站浏览数量
		ApplyGrade int64     `db:"apply_grade"` // 已认证数
		CreateTime time.Time `db:"create_time"`
		UpdateTime time.Time `db:"update_time"`
	}
)

func newDayWebInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultDayWebInfoModel {
	return &defaultDayWebInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`day_web_info`",
	}
}

func (m *defaultDayWebInfoModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	cloudcompDayWebInfoDateKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoDatePrefix, data.Date)
	cloudcompDayWebInfoIdKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, cloudcompDayWebInfoDateKey, cloudcompDayWebInfoIdKey)
	return err
}

func (m *defaultDayWebInfoModel) FindOne(ctx context.Context, id int64) (*DayWebInfo, error) {
	cloudcompDayWebInfoIdKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoIdPrefix, id)
	var resp DayWebInfo
	err := m.QueryRowCtx(ctx, &resp, cloudcompDayWebInfoIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dayWebInfoRows, m.table)
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

func (m *defaultDayWebInfoModel) FindOneByDate(ctx context.Context, date string) (*DayWebInfo, error) {

	cloudcompDayWebInfoDateKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoDatePrefix, date)
	var resp DayWebInfo
	err := m.QueryRowIndexCtx(ctx, &resp, cloudcompDayWebInfoDateKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `date` = '%s' limit 1", dayWebInfoRows, m.table,date)
	
		if err := conn.QueryRowCtx(ctx, &resp, query); err != nil {

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

func (m *defaultDayWebInfoModel) Insert(ctx context.Context, data *DayWebInfo) (sql.Result, error) {
	cloudcompDayWebInfoDateKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoDatePrefix, data.Date)
	cloudcompDayWebInfoIdKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, dayWebInfoRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.Date, data.Login, data.Lookup, data.ApplyGrade)
	}, cloudcompDayWebInfoDateKey, cloudcompDayWebInfoIdKey)
	return ret, err
}

func (m *defaultDayWebInfoModel) Update(ctx context.Context, newData *DayWebInfo) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	cloudcompDayWebInfoDateKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoDatePrefix, data.Date)
	cloudcompDayWebInfoIdKey := fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoIdPrefix, data.Id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, dayWebInfoRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Date, newData.Login, newData.Lookup, newData.ApplyGrade, newData.Id)
	}, cloudcompDayWebInfoDateKey, cloudcompDayWebInfoIdKey)
	return err
}

func (m *defaultDayWebInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCloudcompDayWebInfoIdPrefix, primary)
}

func (m *defaultDayWebInfoModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", dayWebInfoRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultDayWebInfoModel) tableName() string {
	return m.table
}