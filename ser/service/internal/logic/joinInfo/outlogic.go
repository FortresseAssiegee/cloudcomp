package joinInfo

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

func (l *OutLogic) Out(req *types.OutJoinReq) (resp *types.OutJoinResp, err error) {
	// todo: add your logic here and delete this line
	err = l.svcCtx.JoinInfoModel.Delete(l.ctx, req.JoinId)

	if err != nil {
		fmt.Printf("Join活动退出失败，err:%s\n", err)
		return &types.OutJoinResp{
			State:   "failed",
			Commont: "退出失败",
		}, err
	}

	return &types.OutJoinResp{
		State:   "win",
		Commont: "退出成功",
	}, nil
}
