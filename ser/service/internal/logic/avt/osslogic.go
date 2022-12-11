package avt

import (
	"context"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OssLogic {
	return &OssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OssLogic) Oss(req *types.AvtOssReq) (resp *types.AvtOssResp, err error) {
	// todo: add your logic here and delete this line

	return &types.AvtOssResp{
		State:           "win",
		Commont:         "获取成功",
		Region:          "oss-cn-beijing",
		AccessKeyId:     "LTAI5tGBLNYrTJgL7BjEbzJG",
		AccessKeySecret: "8zANAsauIOEX82aEuTe2I8M9LPgm8b",
		Bucket:          "cloud-compe",
		DirPath:         "/avt",
		// region: "oss-cn-chengdu", //地域（在创建 Bucket 的时候指定的中心位置），这里可能不知道具体地域怎么填其实就是 oss-cn-中心位置 ，例：region:'oss-cn-chengdu'，chengdu则是创建bucket是指定的位置成都。
		// accessKeyId: "LTAI5t6eny3M6XmC2m8CfmrK", //阿里云产品的通用id
		// accessKeySecret: "VourrG2Ay39NZ1k63ryPHzoMj9OvCl", //密钥
		// bucket: "cloudcompali", //OSS 存储区域名
	}, nil
}
