package joinUnitInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelExamLogic {
	return &DelExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelExamLogic) DelExam(req *types.DelExamReq) (resp *types.DelExamResp, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.JoinUnitInfoModel.Delete(l.ctx, req.ExamId)
	if err != nil {
		fmt.Printf("JoinUnitInfoModel Delete err:%s\n", err)
		return &types.DelExamResp{
			State:   "win",
			Commont: "设置失败",
		}, err
	}

	
	return &types.DelExamResp{
		State:   "win",
		Commont: "设置成功",
	}, nil
}
