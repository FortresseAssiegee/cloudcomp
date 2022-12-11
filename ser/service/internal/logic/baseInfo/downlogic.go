package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DownLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDownLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DownLogic {
	return &DownLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *DownLogic) Down(req *types.DownReq) (resp *types.DownResp, err error) {
	// todo: add your logic here and delete this line
	actvRes, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	if err != nil {
		fmt.Printf("没有%d,活动,err:%s\n", req.ActvId, err)
		return &types.DownResp{
			Commont: "没有该活动",
			State:   "failed",
		}, err
	}
	newActv := actvRes
	actvRes.StateCode = 10

	err = l.svcCtx.CompeInfoModel.Update(l.ctx, newActv)

	return &types.DownResp{
		Commont: "活动已下架",
		State:   "win",
	}, err
}
