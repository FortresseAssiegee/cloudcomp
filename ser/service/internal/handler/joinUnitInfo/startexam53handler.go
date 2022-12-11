package joinUnitInfo

import (
	"net/http"

	"cldcmp/service/internal/logic/joinUnitInfo"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func StartExam53Handler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.StartExam53Req
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := joinUnitInfo.NewStartExam53Logic(r.Context(), svcCtx)
		resp, err := l.StartExam53(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
