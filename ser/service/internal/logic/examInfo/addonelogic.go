package examInfo

import (
	"context"
	"fmt"
	"strconv"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOneLogic {
	return &AddOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddOneLogic) AddOne(req *types.ExamInfoAddOneReq) (resp *types.ExamInfoAddOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("AddOne", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoAddOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoAddOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	// 查询单元是否存在
	resUnit, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	fmt.Println("该活动的信息", resUnit)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.ExamInfoAddOneResp{
				Commont: "单元不存在",
				State:   "failed",
			}, err
		}
		return &types.ExamInfoAddOneResp{
			Commont: "单元错误",
			State:   "failed",
		}, err
	}

	li := model.ExamInfo{
		UnitId: req.UnitId,
		Title:  req.QuestOneList.Title,
		Quest:  req.QuestOneList.Quest,
		Answer: req.QuestOneList.Answer,
		Img:    req.QuestOneList.Img,
		Method: req.QuestOneList.Method,

		During: 60, //60s,答这道题的时间
		Score:  1,  //分值
	}

	res, err := l.svcCtx.ExamInfoModel.Insert(l.ctx, &li)

	if err != nil {
		fmt.Printf("添加失败")
	}
	fmt.Printf("添加后的结果%+v\n", res)

	getId, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("添加失败")
	}
	fmt.Printf("添加后的结果%+v\n", res)

	str := strconv.FormatInt(req.UserId, 10) + strconv.FormatInt(req.ActvId, 10) + strconv.FormatInt(req.UnitId, 10) + strconv.FormatInt(getId, 10)
	li.Img = str
	li.ExamId = getId
	fmt.Printf("li%+v\nstr%s\n", li, str)

	errUpdate := l.svcCtx.ExamInfoModel.Update(l.ctx, &li)
	if errUpdate != nil {
		return &types.ExamInfoAddOneResp{
			Commont: "修改错误",
			State:   "failed",
		}, err
	}
	fmt.Println("errUpdate", errUpdate)

	return &types.ExamInfoAddOneResp{
		Commont: "添加成功",
		State:   "win",
	}, err

}
