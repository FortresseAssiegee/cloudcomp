package joinInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckJoinLogic {
	return &CheckJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *CheckJoinLogic) CheckJoin(req *types.CheckJoinReq) (resp *types.CheckJoinResp, err error) {
	// todo: add your logic here and delete this line

	fmt.Printf("CheckJoin req:%+v\n", req)
	id := 0
	res, err := l.svcCtx.JoinInfoModel.FindJoinByThreeId(req.ActvId, req.UserCreateId, req.UserJoinId)
	if len(res) == 0 {

		fmt.Printf("未参加")
		return &types.CheckJoinResp{
			State:   "win",
			Commont: "获取成功",
			JoinId:  int64(id),
		}, nil

	}
	return &types.CheckJoinResp{
		State:   "win",
		Commont: "获取成功",
		JoinId:  res[0].JoinId,
	}, nil
}
