package infoTree

import (
	"context"
	"fmt"

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


func (l *AddOneLogic) AddOne(req *types.InfoTreeAddOneReq) (resp *types.InfoTreeAddOneResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("AddOne", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeAddOneResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeAddOneResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeAddOneResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeAddOneResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}
	if req.ParentId != 0 {
		// 找到父节点
		pInfo, err := l.svcCtx.InfoTreeModel.FindOne(l.ctx, req.ParentId)

		// 将父节点content赋值给子节点
		li := model.InfoTree{
			ActvId:   req.ActvId,
			ParentId: req.ParentId,
			Title:    req.Title,
			Content:  pInfo.Content,
		}
		res, err := l.svcCtx.InfoTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("添加失败")
		}
		fmt.Printf("添加后的结果%+v\n", res)

		// 然后将父节点内容清空
		pInfo.Content = ""
		errParent := l.svcCtx.InfoTreeModel.Update(l.ctx, pInfo)
		if errParent != nil {
			return &types.InfoTreeAddOneResp{
				Commont: "修改错误",
				State:   "failed",
			}, err
		}
	} else {
		li := model.InfoTree{
			ActvId:   req.ActvId,
			ParentId: req.ParentId,
			Title:    req.Title,
			Content:  "",
		}
		res, err := l.svcCtx.InfoTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("添加失败")
		}
		fmt.Printf("添加后的结果%+v\n", res)
	}

	// resInfoTree ,err :=l.svcCtx.InfoTreeModel.FindOne(l.ctx,req.Id)

	return &types.InfoTreeAddOneResp{
		Commont: "添加成功",
		State:   "win",
	}, err
}
