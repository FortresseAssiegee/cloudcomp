package baseInfo

import (
	"context"
	"fmt"
	"sort"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotTagShowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotTagShowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotTagShowListLogic {
	return &GetHotTagShowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetHotTagShowListLogic) GetHotTagShowList(req *types.HotTagListReq) (resp *types.HotTagListResp, err error) {
	// todo: add your logic here and delete this line
	hotList, err := l.svcCtx.TagTreeModel.FindAllHot(l.ctx)
	if err != nil {
		fmt.Printf("获取失败,err:%s\n", err)
		return nil, err
	}
	list := []types.HotTag{}
	for _, item := range hotList {
		res, err := l.svcCtx.CompeInfoModel.FindAllActvByTag(item)
		sortedList, err := l.GradeSortList(res)
		if err != nil {
			fmt.Printf("获取活动失败,err:%s\n", err)
			return &types.HotTagListResp{
				State:   "failed",
				Commont: "获取热门活动失败",
			}, err
		}
		tagRes, err := l.svcCtx.TagTreeModel.FindOne(l.ctx, item)

		li := types.HotTag{
			Label:    tagRes.Label,
			TagId:    tagRes.TagId,
			Children: sortedList,
		}

		list = append(list, li)

	}

	// sList :=list[req.]

	return &types.HotTagListResp{
		State:      "win",
		Commont:    "获取成功",
		HotTagList: list,
	}, nil
}

// 对函数获取后进行grade排序
func (l *GetHotTagShowListLogic) GradeSortList(res []*model.CompeInfo) ([]types.GradeBaseAll, error) {

	list := []types.GradeBaseAll{}

	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户信息失败,err%s:\n", err)
			return nil, err
		}
		li := types.GradeBaseAll{
			ActvId: item.ActvId,
			CreateUserName: userInfo.UserName,

			Grade:      userInfo.Grade,
			TagCode:    item.TagCode,
			ExamCode:   item.ExamCode,
			PeriodCode: item.PeriodCode,
			AwardCode:  item.AwardCode,
			ActvName:   item.ActvName,
			Intro:      item.Intro,
			Pstr:       item.Pstr,

			ApplyStartTime: item.ApplyStartTime,
			ApplyEndTime:   item.ApplyEndTime,
		}
		list = append(list, li)
	}

	
	sort.Slice(list, func(i, j int) bool {
		return list[i].Grade > list[j].Grade
	})
	return list, nil
}
