package infoTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelLogic) Del(req *types.InfoTreeDelReq) (resp *types.InfoTreeDelResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("InfoTreeDel", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeDelResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeDelResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resActvId, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resActvId)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeDelResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeDelResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	// 查询该项信息是否存在
	resInfo, err := l.svcCtx.InfoTreeModel.FindOne(l.ctx, req.InfoId)
	fmt.Println("该项信息:", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeDelResp{
				Commont: "该项信息不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeDelResp{
			Commont: "该项信息错误",
			State:   "failed",
		}, err
	}
	
	errDelete := l.svcCtx.InfoTreeModel.Delete(l.ctx, req.InfoId)
	if errDelete != nil {
		return &types.InfoTreeDelResp{
			Commont: "删除错误",
			State:   "failed",
		}, err
	}
	return &types.InfoTreeDelResp{
		Commont: "删除成功",
		State:   "win",
	}, err
}
