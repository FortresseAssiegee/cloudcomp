package grade

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPassGradeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetPassGradeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPassGradeInfoLogic {
	return &GetPassGradeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetPassGradeInfoLogic) GetPassGradeInfo(req *types.GetPassGradeReq) (resp *types.GetPassGradeResp, err error) {
	// todo: add your logic here and delete this line
	var gCode int64
	gCode = 60
	list, err := l.svcCtx.GradeInfoModel.FindAllGradeByCode(l.ctx, gCode, false)
	if err != nil {
		fmt.Println("获取GradeInfoModel失败")
		return &types.GetPassGradeResp{
			State:   "failed",
			Commont: "获取失败",
		}, nil
	}
	giList := []types.GradeInfo{}

	for _, item := range list {
		res, err := l.svcCtx.UserModel.FindOne(l.ctx, item.UserId)
		if err != nil {
			fmt.Printf("%d用户不存在\n", item.UserId)
			return nil, err
		}
		fmt.Printf("用户信息:%+v\n用户名:%s\n", res, res.UserName)
		li := types.GradeInfo{
			GradeId:    item.GradeId,
			UserId:     item.UserId,
			UserName:   res.UserName,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			ApplyIntro: item.ApplyIntro,
			ReplyIntro: item.ReplyIntro,
			GradeCode:  item.GradeCode,
			UserTel:    res.Tel,
			
		}
		giList = append(giList, li)
	}

	return &types.GetPassGradeResp{
		State:         "win",
		Commont:       "获取成功",
		GradeInfoList: giList,
	}, nil
}
