package grade

import (
	"net/http"

	"cldcmp/service/internal/logic/grade"
	"cldcmp/service/internal/svc"
	"cldcmp/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"

)

func ReplyGradeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReplyGradeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := grade.NewReplyGradeLogic(r.Context(), svcCtx)
		resp, err := l.ReplyGrade(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
