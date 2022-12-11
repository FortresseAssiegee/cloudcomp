package login

import (
	"context"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPwdLogic {
	return &ForgetPwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPwdLogic) ForgetPwd(req *types.UserForgetPwdReq) (resp *types.UserForgetPwdResp, err error) {
	// todo: add your logic here and delete this line

	return
}
