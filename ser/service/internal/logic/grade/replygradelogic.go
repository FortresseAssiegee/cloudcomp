package grade

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReplyGradeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReplyGradeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReplyGradeLogic {
	return &ReplyGradeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *ReplyGradeLogic) ReplyGrade(req *types.ReplyGradeReq) (resp *types.ReplyGradeResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.GradeInfoModel.FindOne(l.ctx, req.GradeId)
	res.ReplyIntro = req.ReplyIntro
	res.EndTime = req.EndTime
	res.GradeCode = req.GradeCode
	if req.GradeCode == int64(61) {
		userInfo, err := l.svcCtx.UserModel.FindOne(l.ctx, res.UserId)
		if err != nil {
			fmt.Printf("获取用户信息失败,err:%s\n", err)
			return &types.ReplyGradeResp{
				State:   "failed",
				Commont: "回复失败",
			}, err
		}
		userInfo.Grade = req.Garde
		err = l.svcCtx.UserModel.Update(l.ctx, userInfo)
		if err != nil {
			fmt.Printf("修改等级失败,err:%s\n", err)
			return &types.ReplyGradeResp{
				State:   "failed",
				Commont: "回复失败",
			}, err
		}
	}

	if err != nil {
		fmt.Printf("GradeInfoModel查找失败，err:%s\n", err)
		return &types.ReplyGradeResp{
			State:   "failed",
			Commont: "回复失败",
		}, nil
	}
	err = l.svcCtx.GradeInfoModel.Update(l.ctx, res)
	if err != nil {
		fmt.Printf("GradeInfoModel添加失败,err:%s\n", err)
		return &types.ReplyGradeResp{
			State:   "failed",
			Commont: "回复失败",
		}, nil

	}

	return &types.ReplyGradeResp{
		State:   "win",
		Commont: "回复成功",
	}, nil

}
