package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TagTreeModel = (*customTagTreeModel)(nil)

type (
	// TagTreeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTagTreeModel.
	TagTreeModel interface {
		tagTreeModel
	}

	customTagTreeModel struct {
		*defaultTagTreeModel
	}
)

// NewTagTreeModel returns a model for the database table.
func NewTagTreeModel(conn sqlx.SqlConn, c cache.CacheConf) TagTreeModel {
	return &customTagTreeModel{
		defaultTagTreeModel: newTagTreeModel(conn, c),
	}
}
