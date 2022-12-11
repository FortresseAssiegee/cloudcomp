package unitTree

import (
	"context"
	"fmt"

	"cldcmp/common/diyTools"
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
func (l *AddOneLogic) AddOne(req *types.UnitTreeAddOneReq) (resp *types.UnitTreeAddOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("AddOne", req)

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.UnitTreeAddOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.UnitTreeAddOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	if req.ParentId != 0 {
		// 找到父节点
		pUnit, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.ParentId)
		if err != nil {
			fmt.Printf("未找到该内容")
		}

		// 将父节点content赋值给子节点
		li := model.UnitTree{
			ActvId:   req.ActvId,
			ParentId: req.ParentId,
			Title:    req.Title,
			Tips:     pUnit.Tips,

			During:    pUnit.During,
			Face:      pUnit.Face,
			Count:     pUnit.Count,
			ExamCode:  pUnit.ExamCode,
			Model:     pUnit.Model,
			StartTime: pUnit.StartTime,
			EndTime:   pUnit.EndTime,
		}
		res, err := l.svcCtx.UnitTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("Insert 添加失败,err:%s\n", err)
		}
		fmt.Printf("添加后的结果%+v\n", res)
		unitId, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("LastInsertId 添加失败,err:%s\n", err)
		}
		resUnit, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, unitId)
		if err != nil {
			fmt.Printf("FindOne 添加失败,err:%s\n", err)
		}
		resUnit.FileName = fmt.Sprintf("%d%d%d", req.UserId, req.ActvId, unitId)
		err = l.svcCtx.UnitTreeModel.Update(l.ctx, resUnit)
		if err != nil {
			fmt.Printf("Update 添加失败,err:%s\n", err)
		}

		// 然后将父节点内容清空
		pUnit.Tips = ""
		pUnit.Model = 0

		errParent := l.svcCtx.UnitTreeModel.Update(l.ctx, pUnit)
		if errParent != nil {
			return &types.UnitTreeAddOneResp{
				Commont: "修改错误",
				State:   "failed",
			}, err
		}
	} else {

		sTime, eTime := diyTools.GetDuring(60 * 24)
		li := model.UnitTree{
			ActvId:    req.ActvId,
			ParentId:  req.ParentId,
			Title:     req.Title,
			Model:     int64(70),
			StartTime: sTime,
			EndTime:   eTime,
		}
		res, err := l.svcCtx.UnitTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("添加失败,err:%s\n", err)
		}
		unitId, _ := res.LastInsertId()
		resUnit, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, unitId)
		if err != nil {
			fmt.Printf("FindOne 添加失败,err:%s\n", err)
		}
		resUnit.FileName = fmt.Sprintf("%d%d%d", req.UserId, req.ActvId, unitId)
		fmt.Printf("-------------------------------------------\nfileName:%s\n-----------------", resUnit.FileName)
		// resUnit.FileName = string(req.UserId) + string(req.ActvId) + string(unitId)
		err = l.svcCtx.UnitTreeModel.Update(l.ctx, resUnit)
		if err != nil {
			fmt.Printf("Update 添加失败,err:%s\n", err)
		}
	}

	return &types.UnitTreeAddOneResp{
		Commont: "添加成功",
		State:   "win",
	}, err
}
