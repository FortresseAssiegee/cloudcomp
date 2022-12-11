package infoTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditOneLogic {
	return &EditOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditOneLogic) EditOne(req *types.InfoTreeEditOneReq) (resp *types.InfoTreeEditOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeEditOneResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeEditOneResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeEditOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeEditOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	resInfoTree, err := l.svcCtx.InfoTreeModel.FindOne(l.ctx, req.Id)
	fmt.Println("该活动的信息", resInfoTree)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeEditOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeEditOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	
	li := model.InfoTree{
		InfoId:   resInfoTree.InfoId,
		ActvId:   resInfoTree.ActvId,
		ParentId: resInfoTree.ParentId,
		Title:    req.Title,
		Content:  req.Content,
	}
	errDelete := l.svcCtx.InfoTreeModel.Update(l.ctx, &li)
	if errDelete != nil {
		return &types.InfoTreeEditOneResp{
			Commont: "修改错误",
			State:   "failed",
		}, err
	}
	return &types.InfoTreeEditOneResp{
		Commont: "修改成功",
		State:   "win",
	}, err

}
