package baseInfo

import (
	"net/http"

	"cldcmp/service/internal/logic/baseInfo"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"


)

func AddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BaseInfoAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := baseInfo.NewAddLogic(r.Context(), svcCtx)
		resp, err := l.Add(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
