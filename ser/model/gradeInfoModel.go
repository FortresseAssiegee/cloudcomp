package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ GradeInfoModel = (*customGradeInfoModel)(nil)

type (
	// GradeInfoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGradeInfoModel.
	GradeInfoModel interface {
		gradeInfoModel
	}

	customGradeInfoModel struct {
		*defaultGradeInfoModel
	}
)

// NewGradeInfoModel returns a model for the database table.
func NewGradeInfoModel(conn sqlx.SqlConn, c cache.CacheConf) GradeInfoModel {
	return &customGradeInfoModel{
		defaultGradeInfoModel: newGradeInfoModel(conn, c),
	}
}
