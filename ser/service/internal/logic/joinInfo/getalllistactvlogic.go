package joinInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllListActvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllListActvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllListActvLogic {
	return &GetAllListActvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetAllListActvLogic) GetAllListActv(req *types.GetAllListActvReq) (resp *types.GetAllListActvResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.JoinInfoModel.FindAllJoinByActvId(req.ActvId)
	if err != nil {
		fmt.Printf("获取失败，err:%s\n", err)
		return &types.GetAllListActvResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}

	list := []types.UserInfo{}
	for _, item := range res {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserJoinId)
		if err != nil {
			fmt.Printf("获取用户信息失败,err:%s\n", err)
			return &types.GetAllListActvResp{
				State:   "failed",
				Commont: "获取失败",
			}, nil
		}
		li := types.UserInfo{
			UserName: userInfo.UserName,
			Avt:      userInfo.Avt,
			Tel:      userInfo.Tel,
			UserId:   userInfo.UserId,
			Grade:    userInfo.Grade,
			Email:    userInfo.Email,
			JoinDate: item.CreateTime.Format("2006-01-02 15:04:05"),
		}

		list = append(list, li)
	}

	// sort.Slice(list, func(i, j int) bool {
	// 	return list[i].Grade > list[j].Grade
	// })

	start, end, getPage, LenPage := diyTools.GetPage(req.Count, req.NowPage, int64(len(list)))
	sList := list[start:end]

	return &types.GetAllListActvResp{
		State:        "win",
		Commont:      "获取成功",
		JoinUserList: sList, 
		PageObj: types.Page{
			LenPage: LenPage,
			Count:   req.Count,
			NowPage: getPage,
		},
	}, nil
}
