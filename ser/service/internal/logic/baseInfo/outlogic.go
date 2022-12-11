package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OutLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOutLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OutLogic {
	return &OutLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *OutLogic) Out(req *types.OutReq) (resp *types.OutResp, err error) {
	// todo: add your logic here and delete this line'
	actvRes, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	if err != nil {
		fmt.Printf("没有%d,活动,err:%s\n", req.ActvId, err)
		return &types.OutResp{
			Commont: "没有该活动",
			State:   "failed",
		}, err
	}
	newActv := actvRes
	actvRes.StateCode = 12
	actvRes.EndTime = req.EndTime

	err = l.svcCtx.CompeInfoModel.Update(l.ctx, newActv)

	return &types.OutResp{
		Commont: "活动已结束",
		State:   "win",
	}, err
}
