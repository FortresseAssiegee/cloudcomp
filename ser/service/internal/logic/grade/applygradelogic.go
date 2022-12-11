package grade

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ApplyGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewApplyGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ApplyGradeLogic {
	return &ApplyGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *ApplyGradeLogic) ApplyGrade(req *types.ApplyGradeReq) (resp *types.ApplyGradeResp, err error) {
	// todo: add your logic here and delete this line
	ApplyGrade := &model.GradeInfo{
		UserId:     req.UserId,
		ApplyIntro: req.ApplyIntro,
		StartTime:  req.StartTime,
		GradeCode:  60,
	}
	_, err = l.svcCtx.GradeInfoModel.Insert(l.ctx, ApplyGrade)
	if err != nil {
		fmt.Println("GradeInfoModel添加失败！")
		return &types.ApplyGradeResp{
			State:   "failed",
			Commont: "申请失败",
		}, nil

	}

	return &types.ApplyGradeResp{
		State:   "win",
		Commont: "申请成功，正在审核中，请耐心等待。",
	}, nil

}
