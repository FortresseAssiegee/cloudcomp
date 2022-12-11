package unitTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get(req *types.UnitTreeGetReq) (resp *types.UnitTreeGetResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoDel", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeGetResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeGetResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	treeRows, err := l.svcCtx.UnitTreeModel.FindAllActvByActvId(req.ActvId)
	if err == model.ErrNotFound {
		fmt.Println("没有创建的活动")
		return &types.UnitTreeGetResp{
			Commont: "未创建相关活动",
			State:   "failed",
		}, err
	}
	UnitTree := l.listToTree(treeRows, 0)

	return &types.UnitTreeGetResp{
		Commont:  "获取成功",
		State:    "win",
		UnitTree: UnitTree,
	}, nil
}

// infoTree 是list   treelist 是tree
func (l *GetLogic) listToTree(lis []*model.UnitTree, pid int64) []types.UnitLeafList {
	var treeList []types.UnitLeafList
	// 根据list中pid来获取元素
	for _, item := range lis {
		if pid == item.ParentId {
			node := types.UnitLeafList{
				Id:    item.UnitId,
				Title: item.Title,
				Model: item.Model,

				Child: []types.UnitLeafList{},
			}
			treeList = append(treeList, node)
		}
	}
	for i := 0; i < len(treeList); i++ {
		treeList[i].Child = append(treeList[i].Child, l.listToTree(lis, treeList[i].Id)...)
	}

	return treeList

}
