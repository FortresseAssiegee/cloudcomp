package infoTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *EditLogic) Edit(req *types.InfoTreeEditReq) (resp *types.InfoTreeEditResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("InfoTreeEdit", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeEditResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeEditResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeEditResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeEditResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	treeRows, err := l.svcCtx.InfoTreeModel.FindAllActvByActvId(req.ActvId)
	if err == model.ErrNotFound {
		fmt.Println("没有创建的活动")
		return &types.InfoTreeEditResp{
			Commont: "未创建相关活动",
			State:   "failed",
		}, err
	}
	for _, item := range treeRows {
		errDelete := l.svcCtx.InfoTreeModel.Delete(l.ctx, item.InfoId)
		if errDelete != nil {
			return &types.InfoTreeEditResp{
				Commont: "删除错误",
				State:   "failed",
			}, err
		}
	}

	l.treeToList(req.InfoTree, 0, req.ActvId)

	return &types.InfoTreeEditResp{
		Commont: "修改成功",
		State:   "win",
	}, err
}

func (l *EditLogic) treeToList(tres []types.Leaf, paId, ActvId int64) {
	for _, item := range tres {
		li := model.InfoTree{
			ActvId:   ActvId,
			ParentId: paId,
			Title:    item.Title,
			Content:  item.Content,
		}
		res, err := l.svcCtx.InfoTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("添加失败")
		}
		fmt.Printf("添加后的结果%+v\n", res)
		getId, err := res.LastInsertId()
		// list = append(list, li)
		fmt.Println(getId, "的孩子有", len(item.Child))
		if len(item.Child) != 0 {
			// fmt.Println(item.id, "的", list)
			l.treeToList(item.Child, getId, ActvId)
		}
	}
}
