package userInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetInfoLogic {
	return &GetInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetInfoLogic) GetInfo(req *types.GetInfoReq) (resp *types.GetInfoResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("GetInfo搜：", res)
	if err != nil || res.UserName == "" {
		fmt.Println("账号不存在")
		return &types.GetInfoResp{
			State:   "failed",
			Commont: "账号不存在",
		}, err
	}

	createTab, err := l.svcCtx.CompeInfoModel.FindAllActvByUserId(req.UserId)
	if err != nil {
		fmt.Printf("获取CompeInfoModel失败:%s\n", err)
		return &types.GetInfoResp{
			State:   "failed",
			Commont: "获取信息失败",
		}, err
	}
	createNum := len(createTab)
	fmt.Printf("createNum=%d", createNum)

	joinTab, err := l.svcCtx.JoinInfoModel.FindAllJoinByUserJoinId(req.UserId)
	if err != nil {
		fmt.Printf("获取JoinInfoModel失败:%s\n", err)
		return &types.GetInfoResp{
			State:   "failed",
			Commont: "获取信息失败",
		}, err
	}
	joinNum := len(joinTab)
	fmt.Printf("joinNum=%d", joinNum)

	info := types.Info{
		UserName:   res.UserName,
		Avt:        res.Avt,
		Tel:        res.Tel,
		CreateTime: res.CreateTime.String(),
		UserId:     res.UserId,
		Grade:      res.Grade,
		CreateNum:  int64(createNum),
		JoinNum:    int64(joinNum),
		Email:      res.Email,
		Intro:      res.Intro,
	}
	fmt.Printf("getInfo 发送信息:%v\n", info)

	return &types.GetInfoResp{
		State:    "win",
		Commont:  "登陆成功",
		UserInfo: info,
	}, nil
}
