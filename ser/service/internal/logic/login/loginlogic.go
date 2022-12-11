package login

import (
	"context"
	"fmt"
	"time"

	"cldcmp/common/cryptx"
	"cldcmp/common/jwtx"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	fmt.Println("调用Login", req.UserName, req.Pwd)
	if req.UserName == "" {
		fmt.Println("未输入账号名称")
		return &types.UserLoginResp{
			State:   "failed",
			Commont: "未输入账号名称",
		}, nil
	}

	res, err := l.svcCtx.UserModel.FindOneByUserName(l.ctx, req.UserName)
	fmt.Println("搜：", res)
	if err != nil || res.UserName == "" {
		fmt.Println("账号不存在")
		return &types.UserLoginResp{
			State:   "failed",
			Commont: "账号不存在",
		}, err
	}

	// 判断密码是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Pwd)
	if password != res.Pwd {
		// if req.Pwd != res.Pwd {
		fmt.Println("密码错误")
		return &types.UserLoginResp{
			State:   "failed",
			Commont: "密码错误",
		}, err
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.UserId)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		State:        "win",
		Commont:      "登陆成功",
		UserId:       res.UserId,
		Grade:        res.Grade,
	}, nil
}
