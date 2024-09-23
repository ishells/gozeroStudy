package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goZeroShopMall/apps/app/api/internal/logic"
	"goZeroShopMall/apps/app/api/internal/svc"
	"goZeroShopMall/apps/app/api/internal/types"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
