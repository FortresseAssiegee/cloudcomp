package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DefaultInfoModel = (*customDefaultInfoModel)(nil)

type (
	// DefaultInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDefaultInfoModel.
	DefaultInfoModel interface {
		defaultInfoModel
	}

	customDefaultInfoModel struct {
		*defaultDefaultInfoModel
	}
)

// NewDefaultInfoModel returns a model for the database table.
func NewDefaultInfoModel(conn sqlx.SqlConn, c cache.CacheConf) DefaultInfoModel {
	return &customDefaultInfoModel{
		defaultDefaultInfoModel: newDefaultInfoModel(conn, c),
	}
}
