package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ JoinUnitInfoModel = (*customJoinUnitInfoModel)(nil)

type (
	// JoinUnitInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customJoinUnitInfoModel.
	JoinUnitInfoModel interface {
		joinUnitInfoModel
	}

	customJoinUnitInfoModel struct {
		*defaultJoinUnitInfoModel
	}
)

// NewJoinUnitInfoModel returns a model for the database table.
func NewJoinUnitInfoModel(conn sqlx.SqlConn, c cache.CacheConf) JoinUnitInfoModel {
	return &customJoinUnitInfoModel{
		defaultJoinUnitInfoModel: newJoinUnitInfoModel(conn, c),
	}
}
