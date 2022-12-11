package examInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
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


func (l *GetOneLogic) GetOne(req *types.ExamInfoGetOneReq) (resp *types.ExamInfoGetOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	resUnitTree, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	fmt.Println("该活动的信息", resUnitTree)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	resExam, err := l.svcCtx.ExamInfoModel.FindOne(l.ctx, req.ExamId)
	fmt.Println("该题的信息", resExam)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoGetOneResp{
				Commont: "题目不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoGetOneResp{
			Commont: "题目错误",
			State:   "failed",
		}, err
	}

	node := types.QuestOne{
		Title:  resExam.Title,
		Quest:  resExam.Quest,
		Answer: resExam.Answer,
		Img:    resExam.Img,
		Method: resExam.Method,
		ExamId: resExam.ExamId,
		During: resExam.During,
		Score:  resExam.Score,
	}

	return &types.ExamInfoGetOneResp{
		Commont:      "修改成功",
		State:        "win",
		QuestOneList: node,
	}, err
}
