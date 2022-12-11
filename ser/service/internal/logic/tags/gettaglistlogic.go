package tags

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagListLogic {
	return &GetTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetTagListLogic) GetTagList(req *types.GetTagListReq) (resp *types.GetTagListResp, err error) {
	// todo: add your logic here and delete this line
	list, err := l.svcCtx.TagTreeModel.FindAll(l.ctx)
	if err != nil {
		fmt.Printf("获取所有list失败,err:%s\n", err)
		return &types.GetTagListResp{
			State:   "failed",
			Commont: "获取失败",
		}, err
	}
	newList := l.listToTree(list, 0)

	return &types.GetTagListResp{
		State:       "win",
		Commont:     "获取成功",
		TagInfoList: newList,
	}, nil

}

func (l *GetTagListLogic) listToTree(lis []*model.TagTree, pid int64) []types.Child {
	var Child []types.Child
	// 根据list中pid来获取元素
	for _, item := range lis {
		if pid == item.ParentId {
			node := types.Child{
				TagId:    item.TagId,
				Label:    item.Label,
				Children: []types.Child{},
			}
			Child = append(Child, node)
		}
	}
	for i := 0; i < len(Child); i++ {
		Child[i].Children = append(Child[i].Children, l.listToTree(lis, Child[i].TagId)...)
	}

	return Child

}
