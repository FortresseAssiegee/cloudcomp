package unitTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *GetOneLogic) GetOne(req *types.UnitTreeGetOneReq) (resp *types.UnitTreeGetOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("DelOne", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.UnitTreeGetOneResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeGetOneResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	qtList, err := l.svcCtx.ExamInfoModel.FindAllExamByUnitId(req.UnitId)
	examList := make([]types.QuestTitleOne, 0)
	for _, item := range qtList {
		examList = append(examList, types.QuestTitleOne{
			ExamId: item.ExamId,
			Title:  item.Title,
		})
	}

	resUnitTree, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	fmt.Println("该活动的信息", resUnitTree)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeGetOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeGetOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	node := types.UnitBase{
		Tips: resUnitTree.Tips,

		Title: resUnitTree.Title,

		During:    resUnitTree.During,
		Face:      resUnitTree.Face,
		Count:     resUnitTree.Count,
		ExamCode:  resUnitTree.ExamCode,
		Model:     resUnitTree.Model, //单元模式
		StartTime: resUnitTree.StartTime,
		EndTime:   resUnitTree.EndTime,
		FileName:  resUnitTree.FileName,
	}

	return &types.UnitTreeGetOneResp{
		Commont: "获取成功",
		State:   "win",
		OneNode: node,
		MaxExam: int64(len(qtList)),
	}, err
}
