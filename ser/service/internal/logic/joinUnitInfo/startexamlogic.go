package joinUnitInfo

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"cldcmp/common/diyTools"
	"cldcmp/model"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStartExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartExamLogic {
	return &StartExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}


func (l *StartExamLogic) StartExam(req *types.StartExamReq) (resp *types.StartExamResp, err error) {
	// todo: add your logic here and delete this line
	// todo 检查是否有参加过的

	list, err := l.svcCtx.ExamInfoModel.FindAllExamByUnitId(req.UnitId)
	if err != nil {
		fmt.Printf("获取失败,err:%s\n", err)
		return nil, err
	}
	uRes, err := l.svcCtx.UnitTreeModel.FindOne(l.ctx, req.UnitId)
	if err != nil {
		fmt.Printf("获取失败,err:%s\n", err)
		return nil, err
	}
	qList := []types.TitleExam{}

	len := len(list)
	randList := diyTools.GenerateRandomNumber(0, len, int(uRes.Count))

	//将时间戳设置成种子数
	rand.Seed(time.Now().UnixNano())
	//生成10个0-99之间的随机数
	for _, num := range randList {

		li := types.TitleExam{
			ExamId: list[num].ExamId,
			Title:  list[num].Title,
			Img:    list[num].Img,
			Method: list[num].Method,
			Score:  list[num].Score,
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
		return &types.StartExamResp{
			State:   "failed",
			Commont: "进入考场失败",
		}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("获取添加的Id失败，err:%s\n", err)
		return &types.StartExamResp{
			State:   "failed",
			Commont: "进入考场失败",
		}, err
	}

	return &types.StartExamResp{
		State:       "win",
		Commont:     "进入考场",
		ExamId:      id,
		QuestList:   qList,
		// StartTime:   sTime,
		// MaBeEndTime: eTime,
	}, nil
}
