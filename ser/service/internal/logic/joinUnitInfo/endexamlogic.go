package joinUnitInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EndExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEndExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EndExamLogic {
	return &EndExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EndExamLogic) EndExam(req *types.EndExamReq) (resp *types.EndExamResp, err error) {

	res, err := l.svcCtx.JoinUnitInfoModel.FindOne(l.ctx, req.ExamId)
	if err != nil {
		fmt.Printf("查找考试id失败,err:%s\n", err)
		return &types.EndExamResp{
			State:   "failed",
			Commont: "提交失败",
		}, err
	}

	if res.EndTime != "" {
		return &types.EndExamResp{
			State:   "win",
			Commont: "已提交",
		}, err
	}

	score := float64(0)
	for _, item := range req.AnswerList {
		// 获取题的信息
		eRes, err := l.svcCtx.ExamInfoModel.FindOne(l.ctx, item.ExamId)
		if err != nil {
			fmt.Printf("%d没有该题目,err:%s\n", item.ExamId, err)
		}
		if diyTools.CheckArr(item.Answer, eRes.Answer) {
			score += eRes.Score
		}
	}

	res.Intro = req.RecordInfo.Intro
	res.Illegal = req.RecordInfo.Illegal
	res.IllegalIntro = req.RecordInfo.IllegalIntro
	res.EndTime = req.EndTime
	res.Score = score
	res.StartTime = req.StartTime
	res.EndTime = req.EndTime
	fmt.Printf("考试信息：%+v\n", res)

	err = l.svcCtx.JoinUnitInfoModel.Update(l.ctx, res)

	if err != nil {
		fmt.Printf("更新考试信息失败,err:%s\n", err)
		return &types.EndExamResp{
			State:   "failed",
			Commont: "提交失败",
		}, err
	}
	return &types.EndExamResp{
		State:   "win",
		Commont: "提交成功,考试结束",
	}, nil
}
