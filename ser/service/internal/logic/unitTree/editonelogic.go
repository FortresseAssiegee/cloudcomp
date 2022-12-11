package unitTree

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

func (l *EditOneLogic) EditOne(req *types.UnitTreeEditOneReq) (resp *types.UnitTreeEditOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeEditOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeEditOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	resInfoTree, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	fmt.Println("该活动的信息", resInfoTree)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeEditOneResp{
				Commont: "该节点不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeEditOneResp{
			Commont: "节点错误",
			State:   "failed",
		}, err
	}
	li := model.UnitTree{
		UnitId:   resInfoTree.UnitId,
		ActvId:   resInfoTree.ActvId,
		ParentId: resInfoTree.ParentId,
		FileName: resInfoTree.FileName,

		Title: req.OneNode.Title,
		Tips:  req.OneNode.Tips,

		During:    req.OneNode.During,
		Face:      req.OneNode.Face,
		Count:     req.OneNode.Count,
		ExamCode:  req.OneNode.ExamCode,
		Model:     req.OneNode.Model,
		StartTime: req.OneNode.StartTime,
		EndTime:   req.OneNode.EndTime,
	}

	errDelete := l.svcCtx.UnitTreeModel.Update(l.ctx, &li)
	if errDelete != nil {
		return &types.UnitTreeEditOneResp{
			Commont: "修改错误",
			State:   "failed",
		}, err
	}
	return &types.UnitTreeEditOneResp{
		Commont: "修改成功",
		State:   "win",
	}, err

}
