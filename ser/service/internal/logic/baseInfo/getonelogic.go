package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOneLogic) GetOne(req *types.BaseInfoGetOneReq) (resp *types.BaseInfoGetOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoGet", req)
	// res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	// if err == model.ErrNotFound {
	// 	fmt.Println("已存在的是：", res.UserName)
	// 	return &types.BaseInfoGetOneResp{
	// 		Commont: "用户不存在",
	// 		State:   "failed",
	// 	}, err
	// }

	baseInfoOne, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	if err != nil {
		fmt.Println("没有创建的活动")
		return &types.BaseInfoGetOneResp{
			Commont: "未创建相关活动",
			State:   "failed",
		}, err
	}

	res, err := l.svcCtx.UserModel.FindOne(l.ctx, baseInfoOne.UserId)
	if err != nil {

		fmt.Printf("错误,err:%s\n", err)
		return &types.BaseInfoGetOneResp{
			Commont: "用户不存在",
			State:   "failed",
		}, err
	}

	info := types.BaseAll{
		ActvId:   baseInfoOne.ActvId,
		ActvName: baseInfoOne.ActvName,

		StateCode:  baseInfoOne.StateCode,
		ExamCode:   baseInfoOne.ExamCode,
		TagCode:    baseInfoOne.TagCode,
		PeriodCode: baseInfoOne.PeriodCode,
		AwardCode:  baseInfoOne.AwardCode,

		Intro: baseInfoOne.Intro,
		Pstr:  baseInfoOne.Pstr,

		StartTime:      baseInfoOne.StartTime,
		EndTime:        baseInfoOne.EndTime,
		ApplyStartTime: baseInfoOne.ApplyStartTime,
		ApplyEndTime:   baseInfoOne.ApplyEndTime,
		ExamStartTime:  baseInfoOne.ExamStartTime,
		ExamEndTime:    baseInfoOne.ExamEndTime,

		CreateUserName:  res.UserName,
		UserId:          res.UserId,
		CreateUserGrade: res.Grade,
		CreateUserAvt:   res.Avt,
		CreateUserEmail: baseInfoOne.CreateUserEmail,
		CreateUserTel:   baseInfoOne.CreateUserTel,
	}
	return &types.BaseInfoGetOneResp{
		Commont:  "获取成功",
		State:    "win",
		BaseInfo: info,
	}, err
}
