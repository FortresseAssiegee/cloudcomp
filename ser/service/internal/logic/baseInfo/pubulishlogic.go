package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PubulishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPubulishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PubulishLogic {
	return &PubulishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *PubulishLogic) Pubulish(req *types.PublishReq) (resp *types.PubulishResp, err error) {
	// todo: add your logic here and delete this line'
	actvRes, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	if err != nil {
		fmt.Printf("没有%d,活动,err:%s\n", req.ActvId, err)
		return &types.PubulishResp{
			Commont: "没有该活动",
			State:   "failed",
		}, err
	}
	if actvRes.InfoDegree != 100 {
		return &types.PubulishResp{
			State:   "failed",
			Commont: "请先完善基础信息到100%",
		}, nil
	}
	// treeRows, err := l.svcCtx.UnitTreeModel.FindAllActvByActvId(req.ActvId)
	// if err == model.ErrNotFound {
	// 	fmt.Println("没有创建的活动")
	// 	return &types.PubulishResp{
	// 		Commont: "未创建相关学习单元,",
	// 		State:   "failed",
	// 	}, err
	// }
	// 对于单元信息不完整的时候,也是需要提醒,不准他发布 todo

	newActv := actvRes
	actvRes.StateCode = 11
	actvRes.StartTime = req.StartTime

	err = l.svcCtx.CompeInfoModel.Update(l.ctx, newActv)

	return &types.PubulishResp{
		Commont: "发布成功",
		State:   "win",
	}, err
}
