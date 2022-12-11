package dayActv

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLookupActvLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLookupActvLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLookupActvLogic {
	return &AddLookupActvLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *AddLookupActvLogic) AddLookup(req *types.AddLookupActvReq) (resp *types.AddLookupActvResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.DayAtcvInfoModel.FindOneByDateActv(l.ctx, req.Date, req.ActvId, req.UserId)
	fmt.Printf("-------------------------------------------\nAddLookup res:%+v\nerr:%s\n-------------------------------------------", res, err)
	if len(res) == 0 {
		dateActv := &model.DayAtcvInfo{
			Date:   req.Date,
			ActvId: req.ActvId,
			UserId: req.UserId,
			Lookup: 1,
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

	res[0].Lookup = res[0].Lookup + 1
	err = l.svcCtx.DayAtcvInfoModel.Update(l.ctx, res[0])
	return &types.AddLookupActvResp{
		Commont: "hot",
		State:   "win",
	}, err
}
