package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetAllLogic) GetAll(req *types.BaseInfoGetAllReq) (resp *types.BaseInfoGetAllResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoGet", req)
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	if err == model.ErrNotFound {
		fmt.Println("已存在的是：", res.UserName)
		return &types.BaseInfoGetAllResp{
			Commont: "用户不存在",
			State:   "failed",
		}, err
	}

	compeRows, err := l.svcCtx.CompeInfoModel.FindAllActvByUserId(req.UserId)

	if err == model.ErrNotFound {
		fmt.Println("没有创建的活动")
		return &types.BaseInfoGetAllResp{
			Commont: "未创建相关活动",
			State:   "failed",
		}, err
	}

	compeList := make([]types.BaseTable, 0)
	for _, item := range compeRows {
		compeList = append(compeList, types.BaseTable{
			ActvId:     item.ActvId,
			StateCode:  item.StateCode,
			InfoDegree: item.InfoDegree,
			ActvName:   item.ActvName,

			StartTime: item.StartTime,
			EndTime:   item.EndTime,
		})
	}
	start, end, getPage, LenPage := diyTools.GetPage(req.Count, req.NowPage, int64(len(compeList)))
	sList := compeList[start:end]

	return &types.BaseInfoGetAllResp{
		Commont:  "获取成功",
		State:    "win",
		BaseInfo: sList,
		PageObj: types.Page{
			LenPage: LenPage,
			Count:   req.Count,
			NowPage: getPage,
		},
	}, err

}
