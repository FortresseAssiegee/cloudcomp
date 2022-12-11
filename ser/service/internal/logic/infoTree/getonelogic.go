package infoTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetOneLogic) GetOne(req *types.InfoTreeGetOneReq) (resp *types.InfoTreeGetOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("GetOne", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeGetOneResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeGetOneResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	resInfoTree, err := l.svcCtx.InfoTreeModel.FindOne(l.ctx, req.Id)
	fmt.Println("该活动的信息", resInfoTree)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	node := types.NodeObj{
		Content: resInfoTree.Content,
		Title:   resInfoTree.Title,
	}

	return &types.InfoTreeGetOneResp{
		Commont: "修改成功",
		State:   "win",
		OneNode: node,
	}, err

}
