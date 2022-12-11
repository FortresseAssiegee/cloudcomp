package tags

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditOneTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditOneTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditOneTagLogic {
	return &EditOneTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditOneTagLogic) EditOneTag(req *types.EditOneTagReq) (resp *types.EditOneTagResp, err error) {

	oneRes, err := l.svcCtx.TagTreeModel.FindOne(l.ctx, req.TagId)
	if err != nil {
		fmt.Println("查找标签失败")
		return &types.EditOneTagResp{
			State:   "failed",
			Commont: "修改失败",
		}, nil
	}
	oneRes.Label = req.Label

	err = l.svcCtx.TagTreeModel.Update(l.ctx, oneRes)
	if err != nil {
		fmt.Println("修改标签失败")
		return &types.EditOneTagResp{
			State:   "failed",
			Commont: "修改失败",
		}, nil
	}
	
	return &types.EditOneTagResp{
		State:   "win",
		Commont: "修改成功",
	}, nil
}
