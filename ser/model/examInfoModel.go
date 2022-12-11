package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ExamInfoModel = (*customExamInfoModel)(nil)

type (
	// ExamInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customExamInfoModel.
	ExamInfoModel interface {
		examInfoModel
	}

	customExamInfoModel struct {
		*defaultExamInfoModel
	}
)

// NewExamInfoModel returns a model for the database table.
func NewExamInfoModel(conn sqlx.SqlConn, c cache.CacheConf) ExamInfoModel {
	return &customExamInfoModel{
		defaultExamInfoModel: newExamInfoModel(conn, c),
	}
}
