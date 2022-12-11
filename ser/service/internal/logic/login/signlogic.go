package login

import (
	"context"
	"fmt"
	"time"

	"cldcmp/common/cryptx"
	"cldcmp/common/jwtx"
	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignLogic {
	return &SignLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *SignLogic) Sign(req *types.UserSignReq) (resp *types.UserSignResp, err error) {
	fmt.Println("调用Sign", req)
	res, err := l.svcCtx.UserModel.FindOneByTel(l.ctx, req.Tel)
	if err == nil {
		logx.Info("已存在的是：", res.UserName)
		return &types.UserSignResp{
			State:   "failed",
			Commont: "该用户已存在",
		}, nil
	}
	// dfavt :=l.svcCtx

	if err == model.ErrNotFound {
		newUser := model.UserInfo{
			UserName: req.UserName,
			Avt:      "https://cloud-compe.oss-cn-beijing.aliyuncs.com/default/userAvt/defaultAvt.png",
			Tel:      req.Tel,
			Pwd:      cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.Pwd),
		}
		res, err := l.svcCtx.UserModel.Insert(l.ctx, &newUser)
		if err != nil {
			return &types.UserSignResp{
				State:   "failed",
				Commont: "注册失败",
			}, nil
		}
		newUser.UserId, err = res.LastInsertId()
		if err != nil {
			return &types.UserSignResp{
				State:   "failed",
				Commont: "注册失败",
			}, nil
		}

		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire

		accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, newUser.UserId)
		if err != nil {
			return nil, err
		}

		return &types.UserSignResp{
			AccessToken:  accessToken,
			AccessExpire: now + accessExpire,
			State:        "win",
			Commont:      "注册成功",
			UserId:       newUser.UserId,
		}, nil
	}

	return &types.UserSignResp{
		State:   "failed",
		Commont: "注册失败",
	}, nil
}
