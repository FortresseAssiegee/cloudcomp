package joinUnitInfo

import (
	"context"
	"fmt"
	"strings"

	"cldcmp/common/diyTools"
	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartExam52Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartExam52Logic(ctx context.Context, svcCtx *svc.ServiceContext) *StartExam52Logic {
	return &StartExam52Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *StartExam52Logic) StartExam52(req *types.StartExam52Req) (resp *types.StartExam52Resp, err error) {
	// todo: add your logic here and delete this line
	fmt.Printf("------------------52 调用---------------\n")
	list, err := l.svcCtx.ExamInfoModel.FindAllExamByUnitId(req.UnitId)
	if err != nil {
		fmt.Printf("获取失败,err:%s\n", err)
		return nil, err
	}
	fmt.Printf("list:%+v\n", list)
	uRes, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	if err != nil {
		fmt.Printf("获取失败,err:%s\n", err)
		return nil, err
	}
	fmt.Printf("uRes:%+v\n", uRes)
	qList := []types.TitleExam52{}

	len := len(list)
	randList := diyTools.GenerateRandomNumber(0, len, int(uRes.Count))

	//生成10个0-99之间的随机数
	for _, num := range randList {
		fmt.Printf("%d \n", num)
		li := types.TitleExam52{
			ExamId: list[num].ExamId,
			Title:  list[num].Title,
			Img:    list[num].Img,
			Method: list[num].Method,
			During: list[num].During,
			Score:  list[num].Score,
			Answer: list[num].Answer,
		}
		li.Quest = strings.Split(list[num].Quest, "@.@")
		qList = append(qList, li)
	}
	// sTime, eTime := diyTools.GetDuring(uRes.During)

	// 添加考试记录
	info := &model.JoinUnitInfo{
		ActvId:       req.ActvId,
		UnitId:       req.UnitId,
		UserCreateId: req.UserCreateId,
		UserJoinId:   req.UserJoinId,
		// StartTime:    sTime,
	}
	res, err := l.svcCtx.JoinUnitInfoModel.Insert(l.ctx, info)
	if err != nil {
		fmt.Printf("进入考试失败,err:%s\n", err)
		return &types.StartExam52Resp{
			State:   "failed",
			Commont: "进入考场失败",
		}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("获取添加的Id失败，err:%s\n", err)
		return &types.StartExam52Resp{
			State:   "failed",
			Commont: "进入考场失败",
		}, err
	}

	return &types.StartExam52Resp{
		State:     "win",
		Commont:   "进入考场",
		ExamId:    id,
		QuestList: qList,
		// StartTime:   sTime,
		// MaBeEndTime: eTime,
	}, nil
}
