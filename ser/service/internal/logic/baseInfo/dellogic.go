package baseInfo

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelLogic {
	return &DelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *DelLogic) Del(req *types.BaseInfoDelReq) (resp *types.BaseInfoDelResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoDel", req)


	// 查询活动是否存在
	re, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", re)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.BaseInfoDelResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.BaseInfoDelResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	err = l.svcCtx.CompeInfoModel.Delete(l.ctx, req.ActvId)
	if err != nil {
		return &types.BaseInfoDelResp{
			Commont: "删除错误",
			State:   "failed",
		}, err
	}
	return &types.BaseInfoDelResp{
		Commont: "删除成功",
		State:   "win",
	}, err
}
