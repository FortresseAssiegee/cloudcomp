package unitTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOneLogic {
	return &DelOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *DelOneLogic) DelOne(req *types.UnitTreeDelOneReq) (resp *types.UnitTreeDelOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.UnitTreeDelOneResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeDelOneResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeDelOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeDelOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	
	errDelete := l.svcCtx.UnitTreeModel.Delete(l.ctx, req.UnitId)
	if errDelete != nil {
		return &types.UnitTreeDelOneResp{
			Commont: "删除错误",
			State:   "failed",
		}, err
	}
	return &types.UnitTreeDelOneResp{
		Commont: "添加成功",
		State:   "win",
	}, err
}
