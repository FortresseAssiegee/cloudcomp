package dayActv

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddJoinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJoinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddJoinLogic {
	return &AddJoinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddJoinLogic) AddJoin(req *types.AddJoinReq) (resp *types.AddJoinResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.DayAtcvInfoModel.FindOneByDateActv(l.ctx, req.Date, req.ActvId, req.UserId)
	if err != nil {
		fmt.Printf("AddHot err:%s", err)
		return nil, err
	}
	if len(res) == 0 {
		dateActv := &model.DayAtcvInfo{
			Date:   req.Date,
			ActvId: req.ActvId,
			UserId: req.UserId,
			Join:   1,
		}
		sqlRes, err := l.svcCtx.DayAtcvInfoModel.Insert(l.ctx, dateActv)
		if err != nil {
			fmt.Println("DayAtcvInfoModel数据添加失败！")
			return nil, err
		}

		_, err = sqlRes.LastInsertId()
		if err != nil {
			fmt.Println("DayAtcvInfoModel数据读取失败！")
			return nil, err
		}

	}

	res[0].Join = res[0].Join + 1
	err = l.svcCtx.DayAtcvInfoModel.Update(l.ctx, res[0])
	return &types.AddJoinResp{
		Commont: "hot",
		State:   "win",
	}, err
}
