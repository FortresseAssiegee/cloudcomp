package tags

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelOneTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelOneTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelOneTagLogic {
	return &DelOneTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DelOneTagLogic) DelOneTag(req *types.DelOneTagReq) (resp *types.DelOneTagResp, err error) {

	err = l.svcCtx.TagTreeModel.Delete(l.ctx, req.TagId)
	if err != nil {
		fmt.Println("删除标签失败")
		return &types.DelOneTagResp{
			State:   "failed",
			Commont: "删除失败",
		}, nil
	}

	return &types.DelOneTagResp{
		State:   "win",
		Commont: "删除成功",
	}, nil
}
