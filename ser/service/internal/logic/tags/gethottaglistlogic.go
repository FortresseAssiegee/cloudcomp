package tags

import (
	"context"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetHotTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetHotTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetHotTagListLogic {
	return &GetHotTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetHotTagListLogic) GetHotTagList(req *types.GetHotTagListReq) (resp *types.GetHotTagListResp, err error) {
	// todo: add your logic here and delete this line
	idList, err := l.svcCtx.TagTreeModel.FindAllHot(l.ctx)

	return &types.GetHotTagListResp{
		State:   "win",
		Commont: "获取成功",
		TagList: idList,
	}, nil
}
