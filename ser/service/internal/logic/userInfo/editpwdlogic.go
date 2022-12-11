package userInfo

import (
	"context"
	"fmt"

	"cldcmp/common/cryptx"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditPwdLogic {
	return &EditPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *EditPwdLogic) EditPwd(req *types.EditPwdReq) (resp *types.EditPwdResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("搜：", res)
	if err != nil || res.UserName == "" {
		fmt.Println("账号不存在")
		return &types.EditPwdResp{
			State:   "failed",
			Commont: "账号不存在",
		}, err
	}

	// 判断oldPwd是否正确
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.OldPwd)
	if password != res.Pwd {
		fmt.Println("密码错误")
		return &types.EditPwdResp{
			State:   "failed",
			Commont: "密码错误",
		}, err
	}

	res.Pwd = cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, req.NewPwd)
	// 加密newPwd
	er := l.svcCtx.UserModel.Update(l.ctx, res)
	if er != nil {
		return &types.EditPwdResp{
			State:   "failed",
			Commont: "修改失败",
		}, err
	}

	return &types.EditPwdResp{
		State:   "win",
		Commont: "修改成功",
	}, err
}
