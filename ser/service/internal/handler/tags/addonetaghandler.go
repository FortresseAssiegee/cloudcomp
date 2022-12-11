package tags

import (
	"net/http"

	"cldcmp/service/internal/logic/tags"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func AddOneTagHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddOneTagReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := tags.NewAddOneTagLogic(r.Context(), svcCtx)
		resp, err := l.AddOneTag(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
