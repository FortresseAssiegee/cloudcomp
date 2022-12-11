package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CodeInfoModel = (*customCodeInfoModel)(nil)

type (
	// CodeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCodeInfoModel.
	CodeInfoModel interface {
		codeInfoModel
	}

	customCodeInfoModel struct {
		*defaultCodeInfoModel
	}
)

// NewCodeInfoModel returns a model for the database table.
func NewCodeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) CodeInfoModel {
	return &customCodeInfoModel{
		defaultCodeInfoModel: newCodeInfoModel(conn, c),
	}
}
