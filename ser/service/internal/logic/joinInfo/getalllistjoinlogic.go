package joinInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllListJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllListJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllListJoinLogic {
	return &GetAllListJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllListJoinLogic) GetAllListJoin(req *types.GetAllListJoinReq) (resp *types.GetAllListJoinResp, err error) {

	res, err := l.svcCtx.JoinInfoModel.FindAllJoinByUserJoinId(req.UserId)
	if err != nil {
		fmt.Printf("获取失败，err:%s\n", err)
		return &types.GetAllListJoinResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}

	list := []types.JoinTable{}
	for _, item := range res {
		actvInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, item.ActvId)
		if err != nil {
			fmt.Printf("获取活动信息失败,err:%s\n", err)
			return &types.GetAllListJoinResp{
				State:   "failed",
				Commont: "获取失败",
			}, nil
		}
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, actvInfo.UserId)
		if err != nil {
			fmt.Printf("获取用户信息失败,err:%s\n", err)
			return &types.GetAllListJoinResp{
				State:   "failed",
				Commont: "获取失败",
			}, nil
		}

		li := types.JoinTable{
			ActvId:         actvInfo.ActvId,
			StateCode:      actvInfo.StateCode,
			ActvName:       actvInfo.ActvName,
			ExamStartTime:  actvInfo.ExamStartTime,
			ExamEndTime:    actvInfo.ExamEndTime,
			UserCreateId:   userInfo.UserId,
			UserCreateName: userInfo.UserName,
			// Rank          :nil
		}

		list = append(list, li)
	}

	// sort.Slice(list, func(i, j int) bool {
	// 	return list[i].Grade > list[j].Grade
	// })

	start, end, getPage, LenPage := diyTools.GetPage(req.Count, req.NowPage, int64(len(list)))
	sList := list[start:end]

	return &types.GetAllListJoinResp{
		State:    "win",
		Commont:  "获取成功",
		JoinList: sList,
		PageObj: types.Page{
			LenPage: LenPage,
			Count:   req.Count,
			NowPage: getPage,
		},
	}, nil
}
