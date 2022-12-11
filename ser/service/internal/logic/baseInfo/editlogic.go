package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.BaseInfoEditReq) (resp *types.BaseInfoEditResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoEdit", req)

	// 查询活动是否存在
	re, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.BaseInfo.ActvId)
	fmt.Println("该活动的信息", re)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.BaseInfoEditResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.BaseInfoEditResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	re.ActvName = req.BaseInfo.ActvName
	re.StateCode = req.BaseInfo.StateCode
	re.ExamCode = req.BaseInfo.ExamCode
	re.TagCode = req.BaseInfo.TagCode
	re.AwardCode = req.BaseInfo.AwardCode
	re.PeriodCode = req.BaseInfo.PeriodCode
	re.Intro = req.BaseInfo.Intro
	re.Pstr = req.BaseInfo.Pstr

	re.StartTime = req.BaseInfo.StartTime
	re.EndTime = req.BaseInfo.EndTime
	re.ApplyStartTime = req.BaseInfo.ApplyStartTime
	re.ApplyEndTime = req.BaseInfo.ApplyEndTime
	re.ExamStartTime = req.BaseInfo.ExamStartTime
	re.ExamEndTime = req.BaseInfo.ExamEndTime
	re.CreateUserEmail = req.BaseInfo.CreateUserEmail
	re.CreateUserTel = req.BaseInfo.CreateUserTel

	re.InfoDegree = getInfoDegree(&req.BaseInfo)

	err = l.svcCtx.CompeInfoModel.Update(l.ctx, re)

	if err != nil {
		return &types.BaseInfoEditResp{
			Commont: "活动修改失败",
			State:   "failed",
		}, err
	}

	return &types.BaseInfoEditResp{
		Commont: "修改成功",
		State:   "win",
	}, err
}

func getInfoDegree(req *types.BaseAll) int64 {
	var de int64
	de = 0
	if req.PeriodCode != 0 {
		de += 10
	}
	if req.AwardCode != 0 {
		de += 10
	}
	if req.Intro != "" {
		de += 10
	}
	if req.Pstr != "" {
		de += 10
	}
	if req.TagCode != 0 {
		de += 10
	}
	if req.ExamCode != 0 {
		de += 10
	}
	if req.ActvName != "" {
		de += 10
	}
	if req.StateCode != 0 {
		de += 10
	}
	if req.StartTime != "" {
		de += 10
	}
	if req.EndTime != "" {
		de += 10
	}
	return de
}
