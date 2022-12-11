package joinUnitInfo

import (
	"context"
	"fmt"
	"sort"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRankListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRankListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRankListLogic {
	return &GetRankListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRankListLogic) GetRankList(req *types.GetRankListReq) (resp *types.GetRankListResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("GetRankList req:%d,%d\n", req.UnitId, req.UserCreateId)
	res, err := l.svcCtx.JoinUnitInfoModel.FindJoinByUnitCreateUserId(l.ctx, req.UnitId, req.UserCreateId)
	fmt.Printf("GetRankList res:%v\n", res)
	if err != nil {
		fmt.Printf("获取失败,err:%s", err)
		return &types.GetRankListResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}
	list := []types.RankInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserJoinId)
		if err != nil {
			fmt.Printf("获取用户信息失败,err:%s\n", err)
			return &types.GetRankListResp{
				State:   "failed",
				Commont: "获取失败",
			}, nil
		}
		li := types.RankInfo{
			ExamId:       item.Id,
			ActvId:       item.ActvId,
			UserName:     userInfo.UserName,
			Avt:          userInfo.Avt,
			Tel:          userInfo.Tel,
			UserCreateId: item.UserCreateId,
			UserJoinId:   item.UserJoinId,
			Score:        item.Score,
			Intro:        item.IllegalIntro,
			During:       diyTools.SubTime(item.StartTime, item.EndTime),
		}

		list = append(list, li)
	}
	sort.Slice(list, func(i, j int) bool {
		if list[i].Score > list[j].Score {
			return true
		} else {
			if list[i].During < list[j].During {
				return true
			} else {
				return false
			}
		}
	})

	for n, _ := range list {
		fmt.Printf("%d\n", n+1)
		list[n].Rank = int64(n + 1)

	}
	fmt.Printf("%v\n", list[0])
	start, end, getPage, LenPage := diyTools.GetPage(req.Count, req.NowPage, int64(len(list)))
	sList := list[start:end]
	return &types.GetRankListResp{
		State:   "win",
		Commont: "获取成功",
		List:    sList,
		PageObj: types.Page{
			LenPage: LenPage,
			Count:   req.Count,
			NowPage: getPage,
		},
	}, nil
}
