package examInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetLogic) Get(req *types.ExamInfoGetReq) (resp *types.ExamInfoGetResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoDel", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoGetResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoGetResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	qtList, err := l.svcCtx.ExamInfoModel.FindAllExamByUnitId(req.UnitId)
	examList := make([]types.QuestTitleOne, 0)
	for _, item := range qtList {
		examList = append(examList, types.QuestTitleOne{
			ExamId: item.ExamId,
			Title:  item.Title,
		})
	}

	return &types.ExamInfoGetResp{
		Commont:        "获取成功",
		State:          "win",
		QuestTitleList: examList,
	}, err
}
