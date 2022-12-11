package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DayAtcvInfoModel = (*customDayAtcvInfoModel)(nil)

type (
	// DayAtcvInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDayAtcvInfoModel.
	DayAtcvInfoModel interface {
		dayAtcvInfoModel
	}

	customDayAtcvInfoModel struct {
		*defaultDayAtcvInfoModel
	}
)

// NewDayAtcvInfoModel returns a model for the database table.
func NewDayAtcvInfoModel(conn sqlx.SqlConn, c cache.CacheConf) DayAtcvInfoModel {
	return &customDayAtcvInfoModel{
		defaultDayAtcvInfoModel: newDayAtcvInfoModel(conn, c),
	}
}
