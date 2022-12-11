package userInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditInfoLogic {
	return &EditInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *EditInfoLogic) EditInfo(req *types.EditInfoReq) (resp *types.EditInfoResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("搜：", res)
	if err != nil || res.UserName == "" {
		fmt.Println("账号不存在")
		return &types.EditInfoResp{
			State:   "failed",
			Commont: "账号不存在",
		}, err
	}
	res.UserName = req.UserName

	res.Avt = req.Avt
	res.Intro = req.Intro
	res.Email = req.Email

	er := l.svcCtx.UserModel.Update(l.ctx, res)
	if er != nil {
		fmt.Printf("UserModel.Update，err:%s\n", er)
		return &types.EditInfoResp{
			State:   "failed",
			Commont: "修改失败",
		}, err
	}

	return &types.EditInfoResp{
		State:   "win",
		Commont: "修改成功",
	}, err
}
