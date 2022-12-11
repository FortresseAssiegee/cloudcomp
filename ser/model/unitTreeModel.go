package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UnitTreeModel = (*customUnitTreeModel)(nil)

type (
	// UnitTreeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUnitTreeModel.
	UnitTreeModel interface {
		unitTreeModel
	}

	customUnitTreeModel struct {
		*defaultUnitTreeModel
	}
)

// NewUnitTreeModel returns a model for the database table.
func NewUnitTreeModel(conn sqlx.SqlConn, c cache.CacheConf) UnitTreeModel {
	return &customUnitTreeModel{
		defaultUnitTreeModel: newUnitTreeModel(conn, c),
	}
}
