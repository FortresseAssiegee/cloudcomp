package joinUnitInfo

import (
	"context"
	"fmt"

	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"

	"github.com/zeromicro/go-zero/core/logx"
)

type SimOssLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSimOssLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SimOssLogic {
	return &SimOssLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SimOssLogic) SimOss(req *types.SimOssReq) (resp *types.SimOssResp, err error) {
	// todo: add your logic here and delete this line
	config := sdk.NewConfig()

	credential := credentials.NewAccessKeyCredential("LTAI5tGBLNYrTJgL7BjEbzJG", "8zANAsauIOEX82aEuTe2I8M9LPgm8b")
	/* use STS Token
	credential := credentials.NewStsTokenCredential("<your-access-key-id>", "<your-access-key-secret>", "<your-sts-token>")
	*/
	client, err := imm.NewClientWithOptions("cn-beijing", config, credential)
	if err != nil {
		fmt.Printf("client,err:%s\n", err)
		return &types.SimOssResp{
			State:   "failed",
			Commont: "对比失败",
		}, nil
		// panic(err)
	}

	request := imm.CreateCompareImageFacesRequest()

	request.Scheme = "https"

	request.Project = "cloudcompe"
	request.ImageUriA = req.ImageUriA
	request.ImageUriB = req.ImageUriB

	fmt.Printf("simOss req:%+v\n", req)

	response, err := client.CompareImageFaces(request)
	if err != nil {
		fmt.Printf("simOss,err:%s\n", err)
	}
	fmt.Printf("response is %#v\n", response)

	return &types.SimOssResp{
		State:   "win",
		Commont: "",
		SimVal:  response.Similarity,
	}, nil
}
