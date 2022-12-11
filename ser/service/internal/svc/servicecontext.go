package svc

import (
	"cldcmp/model"
	"cldcmp/service/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config           config.Config
	CompeInfoModel   model.CompeInfoModel
	InfoTreeModel    model.InfoTreeModel
	UnitTreeModel    model.UnitTreeModel
	UserModel        model.UserInfoModel
	ExamInfoModel    model.ExamInfoModel
	DefaultInfoModel model.DefaultInfoModel

	JoinInfoModel     model.JoinInfoModel
	JoinUnitInfoModel model.JoinUnitInfoModel

	GradeInfoModel   model.GradeInfoModel
	TagTreeModel     model.TagTreeModel
	DayWebInfoModel  model.DayWebInfoModel
	DayAtcvInfoModel model.DayAtcvInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{

		Config:            c,
		GradeInfoModel:    model.NewGradeInfoModel(conn, c.CacheRedis),
		CompeInfoModel:    model.NewCompeInfoModel(conn, c.CacheRedis),
		InfoTreeModel:     model.NewInfoTreeModel(conn, c.CacheRedis),
		UnitTreeModel:     model.NewUnitTreeModel(conn, c.CacheRedis),
		UserModel:         model.NewUserInfoModel(conn, c.CacheRedis),
		ExamInfoModel:     model.NewExamInfoModel(conn, c.CacheRedis),
		DefaultInfoModel:  model.NewDefaultInfoModel(conn, c.CacheRedis),
		DayAtcvInfoModel:  model.NewDayAtcvInfoModel(conn, c.CacheRedis),
		JoinInfoModel:     model.NewJoinInfoModel(conn, c.CacheRedis),
		JoinUnitInfoModel: model.NewJoinUnitInfoModel(conn, c.CacheRedis),
		TagTreeModel:      model.NewTagTreeModel(conn, c.CacheRedis),

		DayWebInfoModel: model.NewDayWebInfoModel(conn, c.CacheRedis),
	}
}
