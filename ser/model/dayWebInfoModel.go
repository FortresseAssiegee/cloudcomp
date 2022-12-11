package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DayWebInfoModel = (*customDayWebInfoModel)(nil)

type (
	// DayWebInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDayWebInfoModel.
	DayWebInfoModel interface {
		dayWebInfoModel
	}

	customDayWebInfoModel struct {
		*defaultDayWebInfoModel
	}
)

// NewDayWebInfoModel returns a model for the database table.
func NewDayWebInfoModel(conn sqlx.SqlConn, c cache.CacheConf) DayWebInfoModel {
	return &customDayWebInfoModel{
		defaultDayWebInfoModel: newDayWebInfoModel(conn, c),
	}
}
