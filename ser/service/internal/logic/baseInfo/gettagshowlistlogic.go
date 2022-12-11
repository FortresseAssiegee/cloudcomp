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

type GetTagShowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagShowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagShowListLogic {
	return &GetTagShowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetTagShowListLogic) GetTagShowList(req *types.GetTagShowListReq) (resp *types.GetTagShowListResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("tagCode", req.TagCode)
	res, err := l.svcCtx.CompeInfoModel.FindAllActvByPeriodAwardTagExam(req.PeriodCode, req.AwardCode, req.TagCode, req.ExamCode)
	if err != nil {
		fmt.Printf("查询失败,err:%s\n", err)
		return &types.GetTagShowListResp{
			State:   "failed",
			Commont: "查询失败",
		}, err

	}

	list, err := l.GradeSortList(res)
	start, end, getPage, LenPage := diyTools.GetPage(req.Count, req.NowPage, int64(len(list)))
	sList := list[start:end]

	if err != nil {
		fmt.Printf("排序失败,err:%s\n", err)
		return &types.GetTagShowListResp{
			State:   "failed",
			Commont: "查询失败",
		}, err
	}
	return &types.GetTagShowListResp{
		State:    "win",
		Commont:  "查询成功",
		BaseInfo: sList,
		PageObj: types.Page{
			LenPage: LenPage,
			Count:   req.Count,
			NowPage: getPage,
		},
	}, nil
}

// 对函数获取后进行grade排序
func (l *GetTagShowListLogic) GradeSortList(res []*model.CompeInfo) ([]types.GradeBaseAll, error) {

	list := []types.GradeBaseAll{}

	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("查找用户信息失败,err%s:\n", err)
			return nil, err
		}
		li := types.GradeBaseAll{
			ActvId: item.ActvId,

			Grade:          userInfo.Grade,
			UserId:         userInfo.UserId,
			CreateUserName: userInfo.UserName,

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
