package baseInfo

import (
	"net/http"

	"cldcmp/service/internal/logic/baseInfo"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func GetHotTagShowListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.HotTagListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := baseInfo.NewGetHotTagShowListLogic(r.Context(), svcCtx)
		resp, err := l.GetHotTagShowList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
