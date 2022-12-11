package grade

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllGradeLogic {
	return &GetAllGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllGradeLogic) GetAllGrade(req *types.GetAllGradeReq) (resp *types.GetAllGradeResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("req num", req.Num)
	res, err := l.svcCtx.UserModel.FindNumByGrade(l.ctx, req.Num)
	fmt.Println("GetAllGrade res:", res)
	if err != nil {
		fmt.Printf("______________________________________GetAllGrade err:%s_______________________________________\n", err)
		return &types.GetAllGradeResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}
	list := []types.UserInfoGrade{}

	for i := 0; i < len(res); i++ {
		li := types.UserInfoGrade{

			UserName: res[i].UserName,
			Avt:      res[i].Avt,
			Intro:    res[i].Intro,
		}
		list = append(list, li)
	}

	return &types.GetAllGradeResp{
		State:    "win",
		Commont:  "获取成功",
		UserList: list,
	}, nil
}
