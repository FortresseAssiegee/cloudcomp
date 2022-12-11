package joinInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinLogic {
	return &JoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinLogic) Join(req *types.JoinReq) (resp *types.JoinResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("Join req:%+v\n", req)
	findList, err := l.svcCtx.JoinInfoModel.FindJoinByThreeId(req.ActvId, req.UserCreateId, req.UserJoinId)
	if len(findList) != 0 {
		return &types.JoinResp{
			State:   "filed",
			Commont: "已收藏该活动",
		}, err
	}

	res, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	if err != nil {
		fmt.Printf("Join查询活动失败，err:%s\n", err)
		return &types.JoinResp{
			State:   "failed",
			Commont: "加入失败",
		}, err
	}
	info := &model.JoinInfo{
		UserJoinId:   req.UserJoinId,
		ActvId:       req.ActvId,
		UserCreateId: res.UserId,
	}
	_, err = l.svcCtx.JoinInfoModel.Insert(l.ctx, info)
	if err != nil {
		fmt.Printf("Join活动加入失败，err:%s\n", err)
		return &types.JoinResp{
			State:   "failed",
			Commont: "加入失败",
		}, err
	}

	return &types.JoinResp{
		State:   "win",
		Commont: "加入成功",
	}, nil
}
