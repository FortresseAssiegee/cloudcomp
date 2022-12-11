package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankAllLogic {
	return &GetRankAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankAllLogic) GetRankAll(req *types.GetRankAllReq) (resp *types.GetRankAllResp, err error) {
	// todo: add your logic here and delete this line
	weekStr := diyTools.GetWeek()
	res, err := l.svcCtx.DayAtcvInfoModel.RankActvIdByHot(l.ctx, weekStr)
	if err != nil {
		fmt.Printf("------------------RankActvIdByHot err :%s----------------------\n", err)

		return nil, err
	}
	HotList := []types.RankBaseInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户%d信息失败,err%s:\n", item.UserId, err)
			continue

		}
		actvInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, item.ActvId)
		if err != nil {
			fmt.Printf("查找活动%d信息失败,err%s:\n", item.ActvId, err)
			continue
		}
		li := types.RankBaseInfo{
			ActvId:         item.ActvId,
			UserId:         userInfo.UserId,
			CreateUserName: userInfo.UserName,
			ActvName:       actvInfo.ActvName,
			Pstr:           actvInfo.Pstr,
		}
		HotList = append(HotList, li)
	}

	res, err = l.svcCtx.DayAtcvInfoModel.RankActvIdByLookup(l.ctx, weekStr)
	if err != nil {
		fmt.Printf("------------------RankActvIdByHot err :%s----------------------\n", err)

		return nil, err
	}
	LookupList := []types.RankBaseInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户%d信息失败,err%s:\n", item.UserId, err)
			continue

		}
		actvInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, item.ActvId)
		if err != nil {
			fmt.Printf("查找活动%d信息失败,err%s:\n", item.ActvId, err)
			continue
		}
		li := types.RankBaseInfo{
			ActvId:         item.ActvId,
			UserId:         userInfo.UserId,
			CreateUserName: userInfo.UserName,
			ActvName:       actvInfo.ActvName,
			Pstr:           actvInfo.Pstr,
		}
		LookupList = append(LookupList, li)
	}

	res, err = l.svcCtx.DayAtcvInfoModel.RankUserIdByHot(l.ctx, weekStr)
	if err != nil {
		fmt.Printf("------------------RankActvIdByHot err :%s----------------------\n", err)

		return nil, err
	}
	UserList := []types.RankUserInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户%d信息失败,err%s:\n", item.UserId, err)
			continue

		}
		li := types.RankUserInfo{

			UserId:   userInfo.UserId,
			UserName: userInfo.UserName,

			Avt: userInfo.Avt,
		}
		UserList = append(UserList, li)
	}
	if len(HotList) > 5 {
		HotList = HotList[:5]
	}
	if len(LookupList) > 5 {
		LookupList = LookupList[:5]
	}
	if len(UserList) > 5 {
		UserList = UserList[:5]
	}

	return &types.GetRankAllResp{
		Commont:         "获取成功",
		State:           "win",
		WeekHotRank:     HotList,
		WeekLookUpRank:  LookupList,
		SchoolCompeRank: UserList,
	}, nil
}
