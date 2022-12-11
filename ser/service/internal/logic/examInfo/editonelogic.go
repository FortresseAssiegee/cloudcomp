package examInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditOneLogic {
	return &EditOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *EditOneLogic) EditOne(req *types.ExamInfoEditOneReq) (resp *types.ExamInfoEditOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoEditOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoEditOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	// 查询单元是否存在
	resUnit, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	fmt.Println("该活动的信息", resUnit)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoEditOneResp{
				Commont: "单元不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoEditOneResp{
			Commont: "单元错误",
			State:   "failed",
		}, err
	}

	resExam, err := l.svcCtx.ExamInfoModel.FindOne(l.ctx, req.ExamId)
	fmt.Println("该题的信息", resExam)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoEditOneResp{
				Commont: "题目不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoEditOneResp{
			Commont: "题目错误",
			State:   "failed",
		}, err
	}
	li := model.ExamInfo{
		ExamId: resExam.ExamId,
		UnitId: resExam.UnitId,
		Title:  req.QuestOneList.Title,
		Quest:  req.QuestOneList.Quest,
		Answer: req.QuestOneList.Answer,
		Img:    req.QuestOneList.Img,
		Method: req.QuestOneList.Method,
		During: req.QuestOneList.During,
		Score:  req.QuestOneList.Score,
	}
	errDelete := l.svcCtx.ExamInfoModel.Update(l.ctx, &li)

	if errDelete != nil {
		fmt.Printf("err:%s\n",errDelete)
		return &types.ExamInfoEditOneResp{
			Commont: "修改错误",
			State:   "failed",
		}, err
	}

	return &types.ExamInfoEditOneResp{
		Commont: "修改成功",
		State:   "win",
	}, err
}
