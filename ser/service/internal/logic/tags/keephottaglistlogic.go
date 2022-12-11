package tags

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KeepHotTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeepHotTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KeepHotTagListLogic {
	return &KeepHotTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KeepHotTagListLogic) KeepHotTagList(req *types.KeepHotTagListReq) (resp *types.KeepHotTagListResp, err error) {
	// todo: add your logic here and delete this line
	oldList, err := l.svcCtx.TagTreeModel.FindAllHot(l.ctx)
	if err != nil {

		return &types.KeepHotTagListResp{
			State:   "failed",
			Commont: "修改失败",
		}, nil
	}

	for _, item := range oldList {

		oneTag, err := l.svcCtx.TagTreeModel.FindOne(l.ctx, item)
		if err != nil {
			fmt.Printf("查找%d错误是:%s\n", item, err)
		}
		oneTag.Hot = 0
		err = l.svcCtx.TagTreeModel.Update(l.ctx, oneTag)
		if err != nil {
			fmt.Printf("修改%d错误是:%s\n", item, err)
		}
	}
	
	for _, item := range req.TagList {

		oneTag, err := l.svcCtx.TagTreeModel.FindOne(l.ctx, item)
		if err != nil {
			fmt.Printf("查找%d错误是:%s\n", item, err)
		}
		oneTag.Hot = 1
		err = l.svcCtx.TagTreeModel.Update(l.ctx, oneTag)
		if err != nil {
			fmt.Printf("修改%d错误是:%s\n", item, err)
		}
	}

	return &types.KeepHotTagListResp{
		State:   "win",
		Commont: "保存成功",
	}, nil
}
