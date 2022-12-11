package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JoinInfoModel = (*customJoinInfoModel)(nil)

type (
	// JoinInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJoinInfoModel.
	JoinInfoModel interface {
		joinInfoModel
	}

	customJoinInfoModel struct {
		*defaultJoinInfoModel
	}
)

// NewJoinInfoModel returns a model for the database table.
func NewJoinInfoModel(conn sqlx.SqlConn, c cache.CacheConf) JoinInfoModel {
	return &customJoinInfoModel{
		defaultJoinInfoModel: newJoinInfoModel(conn, c),
	}
}
