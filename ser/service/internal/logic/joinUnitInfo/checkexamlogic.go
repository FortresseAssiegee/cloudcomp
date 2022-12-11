package joinUnitInfo

import (
	"context"
	"fmt"
	"sort"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckExamLogic {
	return &CheckExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckExamLogic) CheckExam(req *types.CheckExamReq) (resp *types.CheckExamResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("req:%+v\n", req)
	res, err := l.svcCtx.JoinUnitInfoModel.FindJoinByUnitJoinUserId(l.ctx, req.UnitId, req.UserJoinId)

	fmt.Printf("FindJoinByUnitCreateUserId err:%#+v\n", len(res))
	if len(res) == 0 {
		// if err == model.ErrNotFound {
		return &types.CheckExamResp{
			State:   "win",
			Commont: "未参加",
		}, nil

		// } else {
		// 	return &types.CheckExamResp{
		// 		State:   "failed",
		// 		Commont: "获取错误",
		// 	}, err
		// }
	}
	obj := types.JoinUnit{
		ExamId:    res[0].Id,
		Score:     res[0].Score,
		StartTime: res[0].StartTime,
		EndTime:   res[0].EndTime,
	}

	List, err := l.svcCtx.JoinUnitInfoModel.FindJoinByUnitId(l.ctx, req.UnitId)
	sort.Slice(List, func(i, j int) bool {
		return List[i].Score > List[j].Score
	})

	for num, item := range List {
		if item.Id == res[0].Id {
			obj.Rank = int64(num + 1)
		}
	}

	return &types.CheckExamResp{
		State:        "win",
		Commont:      "已参加",
		JoinUnitInfo: obj,
	}, err
}
