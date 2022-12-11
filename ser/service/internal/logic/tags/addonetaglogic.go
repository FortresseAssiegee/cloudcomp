package tags

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOneTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOneTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOneTagLogic {
	return &AddOneTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *AddOneTagLogic) AddOneTag(req *types.AddOneTagReq) (resp *types.AddOneTagResp, err error) {
	// todo: add your logic here and delete this line

	new := model.TagTree{
		ParentId: req.ParentId,
		Label:    req.Label,
		Hot:      0, //0 false
	}
	_, err = l.svcCtx.TagTreeModel.Insert(l.ctx, &new)
	if err != nil {
		fmt.Println("添加新标签失败")
		return &types.AddOneTagResp{
			State:   "failed",
			Commont: "添加失败",
		}, nil
	}
	
	return &types.AddOneTagResp{
		State:   "win",
		Commont: "添加成功",
	}, nil
}

