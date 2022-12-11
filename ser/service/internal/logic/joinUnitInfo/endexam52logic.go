package joinUnitInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndExam52Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndExam52Logic(ctx context.Context, svcCtx *svc.ServiceContext) *EndExam52Logic {
	return &EndExam52Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *EndExam52Logic) EndExam52(req *types.EndExam52Req) (resp *types.EndExam52Resp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.JoinUnitInfoModel.FindOne(l.ctx, req.ExamId)
	if err != nil {
		fmt.Printf("查找考试id失败,err:%s\n", err)
		return &types.EndExam52Resp{
			State:   "failed",
			Commont: "提交失败",
		}, err
	}

	if res.EndTime != "" {
		return &types.EndExam52Resp{
			State:   "win",
			Commont: "已提交",
		}, err
	}


	res.Intro = req.RecordInfo.Intro
	res.Illegal = req.RecordInfo.Illegal
	res.IllegalIntro = req.RecordInfo.IllegalIntro
	res.EndTime = req.EndTime
	res.Score = req.Score
	res.StartTime = req.StartTime
	res.EndTime = req.EndTime
	fmt.Printf("考试信息：%+v\n", res)

	err = l.svcCtx.JoinUnitInfoModel.Update(l.ctx, res)

	if err != nil {
		fmt.Printf("更新考试信息失败,err:%s\n", err)
		return &types.EndExam52Resp{
			State:   "failed",
			Commont: "提交失败",
		}, err
	}
	return &types.EndExam52Resp{
		State:   "win",
		Commont: "提交成功,考试结束",
	}, nil
}
