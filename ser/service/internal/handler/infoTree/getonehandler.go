package infoTree

import (
	"net/http"

	"cldcmp/service/internal/logic/infoTree"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"


)

func GetOneHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InfoTreeGetOneReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := infoTree.NewGetOneLogic(r.Context(), svcCtx)
		resp, err := l.GetOne(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
