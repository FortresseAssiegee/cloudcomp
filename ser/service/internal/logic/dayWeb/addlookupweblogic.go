package dayWeb

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLookupWebLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLookupWebLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLookupWebLogic {
	return &AddLookupWebLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}
func (l *AddLookupWebLogic) AddLookup(req *types.AddLookupWebReq) (resp *types.AddLookupWebResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("AddLookupReq req:%+v\n", req)
	res, err := l.svcCtx.DayWebInfoModel.FindOneByDate(l.ctx, req.Date)

	if err == model.ErrNotFound {
		date := &model.DayWebInfo{
			Date: req.Date,
		}
		_, err := l.svcCtx.DayWebInfoModel.Insert(l.ctx, date)
		if err != nil {
			fmt.Println("DayWebInfoModel数据添加失败！")
			return nil, err
		}

		res, err = l.svcCtx.DayWebInfoModel.FindOneByDate(l.ctx, req.Date)
		if err != nil {
			fmt.Println("DayWebInfoModel数据读取失败！")
			return &types.AddLookupWebResp{
				Commont: "不存在该记录",
				State:   "failed",
			}, err
		}
	}

	res.Lookup++
	err = l.svcCtx.DayWebInfoModel.Update(l.ctx, res)
	return &types.AddLookupWebResp{
		Commont: "hot",
		State:   "win",
	}, err
}
