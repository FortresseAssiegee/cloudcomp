package grade

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneGradeLogic {
	return &GetOneGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetOneGradeLogic) GetOneGrade(req *types.GetOneGradeReq) (resp *types.GetOneGradeResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("GetOneGrade userId%d\n", req.UserId)
	list, err := l.svcCtx.GradeInfoModel.FindAllGradeByUserId(l.ctx, req.UserId)
	if err != nil {
		fmt.Println("获取GradeInfoModel失败")
		return &types.GetOneGradeResp{
			State:   "failed",
			Commont: "获取失败",
		}, nil
	}

	giList := []types.GradeInfo{}

	for _, item := range list {
		fmt.Printf("GetOneGrade list:%+v\n", item)
		li := types.GradeInfo{
			GradeId:    item.GradeId,
			UserId:     item.UserId,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			ApplyIntro: item.ApplyIntro,
			ReplyIntro: item.ReplyIntro,
		}
		giList = append(giList, li)
	}

	return &types.GetOneGradeResp{
		State:         "win",
		Commont:       "获取成功",
		GradeInfoList: giList,
	}, nil

}
