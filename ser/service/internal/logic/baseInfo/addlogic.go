package baseInfo

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


func (l *AddLogic) Add(req *types.BaseInfoAddReq) (resp *types.BaseInfoAddResp, err error) {
	// todo: add your logic here and delete this line
	fmt.Println("BaseInfoAdd", req)
	res, err := l.svcCtx.UserModel.FindOne(l.ctx, req.UserId)
	fmt.Println("已存在的是：", res.UserName, res)
	if err == model.ErrNotFound {
		return &types.BaseInfoAddResp{
			Commont: "用户不存在",
			State:   "failed",
		}, err
	}
	newActv := model.CompeInfo{
		UserId:          req.UserId,
		ActvName:        req.ActvName,
		StateCode:       10,
		Pstr:            "https://cloud-compe.oss-cn-beijing.aliyuncs.com/default/pstr/pic2.png",
		HotNum:          0,
		InfoDegree:      10,
		CreateUserTel:   res.Tel,
		CreateUserEmail: res.Email,
	}

	re, err := l.svcCtx.CompeInfoModel.Insert(l.ctx, &newActv)

	if err != nil {
		fmt.Println("添加时的错误原因", err)
		return &types.BaseInfoAddResp{
			Commont: "用户不存在",
			State:   "failed",
		}, nil
	}

	fmt.Println("添加后返回的结果", re)

	// newActv.ActvId, err = re.LastInsertId()

	return &types.BaseInfoAddResp{
		State:   "win",
		Commont: "添加成功",
	}, nil
}
