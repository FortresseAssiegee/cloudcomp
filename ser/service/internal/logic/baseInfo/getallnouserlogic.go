package baseInfo

import (
	"context"
	"fmt"
	"sort"

	"cldcmp/common/diyTools"
	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllNoUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllNoUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllNoUserLogic {
	return &GetAllNoUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetAllNoUserLogic) GetAllNoUser(req *types.GetAllNoUserReq) (resp *types.GetAllNoUserResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.CompeInfoModel.FindAllActvByNoUser()
	if err != nil {
		fmt.Printf("查找失败,err:%s\n", err)
		return &types.GetAllNoUserResp{
			State:   "failed",
			Commont: "获取数据失败",
		}, err
	}

	list, err := l.GradeSortList(res)
	if err != nil {
		fmt.Printf("错误原因,err:%s\n", err)
		return nil, err
	}
	return &types.GetAllNoUserResp{
		State:    "win",
		Commont:  "获取成功",
		BaseInfo: list,
	}, nil
}

// 对函数获取后进行grade排序
func (l *GetAllNoUserLogic) GradeSortList(res []*model.CompeInfo) ([]types.GradeBaseAll, error) {

	list := []types.GradeBaseAll{}

	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户信息失败,err%s:\n", err)
			return nil, err
		}

		li := types.GradeBaseAll{
			ActvId: item.ActvId,

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

	slist := []types.GradeBaseAll{}

	nums := diyTools.GenerateRandomNumber(0, len(list), 5)
	for _, num := range nums {
		slist = append(slist, list[num])
	}
	return slist, nil
}
