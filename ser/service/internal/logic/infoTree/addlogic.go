package infoTree

import (
	"context"
	"fmt"

	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *AddLogic) Add(req *types.InfoTreeAddReq) (resp *types.InfoTreeAddResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("InfoTreeAdd", req)
	resUser, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", resUser.UserName, resUser)
	if err != nil {

		if err == model.ErrNotFound {
			return &types.InfoTreeAddResp{
				Commont: "用户不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeAddResp{
			Commont: "用户非法",
			State:   "failed",
		}, err
	}

	// 查询活动是否存在
	resInfo, err := l.svcCtx.CompeInfoModel.FindOne(l.ctx, req.ActvId)
	fmt.Println("该活动的信息", resInfo)
	if err != nil {
		if err == model.ErrNotFound {
			return &types.InfoTreeAddResp{
				Commont: "活动不存在",
				State:   "failed",
			}, err
		}
		return &types.InfoTreeAddResp{
			Commont: "活动错误",
			State:   "failed",
		}, err
	}

	l.treeToList(req.InfoTree, 0, req.ActvId)

	return &types.InfoTreeAddResp{
		Commont: "添加成功",
		State:   "win",
	}, err
}

func (l *AddLogic) treeToList(tres []types.Leaf, paId, ActvId int64) {
	for _, item := range tres {
		li := model.InfoTree{
			ActvId:   ActvId,
			ParentId: paId,
			Title:    item.Title,
			Content:  item.Content,
		}
		res, err := l.svcCtx.InfoTreeModel.Insert(l.ctx, &li)
		if err != nil {
			fmt.Printf("添加失败,err:%s\n", err)
		}
		fmt.Printf("添加后的结果%+v\n", res)
		getId, err := res.LastInsertId()
		if err != nil {
			fmt.Printf("getId添加失败,err:%s\n", err)
			// return nil, err
		}
		// list = append(list, li)
		fmt.Println(getId, "的孩子有", len(item.Child))
		if len(item.Child) != 0 {
			// fmt.Println(item.id, "的", list)
			l.treeToList(item.Child, getId, ActvId)
		}
	}
}
