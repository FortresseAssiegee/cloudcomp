package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ InfoTreeModel = (*customInfoTreeModel)(nil)

type (
	// InfoTreeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInfoTreeModel.
	InfoTreeModel interface {
		infoTreeModel
	}

	customInfoTreeModel struct {
		*defaultInfoTreeModel
	}
)

// NewInfoTreeModel returns a model for the database table.
func NewInfoTreeModel(conn sqlx.SqlConn, c cache.CacheConf) InfoTreeModel {
	return &customInfoTreeModel{
		defaultInfoTreeModel: newInfoTreeModel(conn, c),
	}
}
