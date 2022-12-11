package joinUnitInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListLogic {
	return &GetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetListLogic) GetList(req *types.GetListReq) (resp *types.GetListResp, err error) {
	res, err := l.svcCtx.JoinUnitInfoModel.FindJoinByUnitCreateUserId(l.ctx,req.UnitId, req.UserCreateId)
	if err != nil {
		fmt.Printf("获取失败,err:%s", err)
		return &types.GetListResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}
	list := []types.JoinUnitInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserJoinId)
		if err != nil {
			fmt.Printf("获取用户信息失败,err:%s\n", err)
			return &types.GetListResp{
				State:   "failed",
				Commont: "获取失败",
			}, nil
		}
		li := types.JoinUnitInfo{
			ExamId:       item.Id,
			UserName:     userInfo.UserName,
			Tel:          userInfo.Tel,
			UserJoinId:   item.UserJoinId,
			Score:        item.Score,
			Intro:        item.Intro,
			During:      diyTools.SubTime(item.StartTime,item.EndTime),
			Illegal:      item.Illegal,
			IllegalIntro: item.IllegalIntro,
		}

		list = append(list, li)
	}
	return &types.GetListResp{
		State:   "win",
		Commont: "获取成功",
		List:    list,
	}, nil
}
