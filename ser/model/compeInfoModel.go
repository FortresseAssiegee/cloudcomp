package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CompeInfoModel = (*customCompeInfoModel)(nil)

type (
	// CompeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCompeInfoModel.
	CompeInfoModel interface {
		compeInfoModel
	}

	customCompeInfoModel struct {
		*defaultCompeInfoModel
	}
)

// NewCompeInfoModel returns a model for the database table.
func NewCompeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) CompeInfoModel {
	return &customCompeInfoModel{
		defaultCompeInfoModel: newCompeInfoModel(conn, c),
	}
}
