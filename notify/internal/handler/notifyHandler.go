package handler

import (
	"cloudflare-pages-hook/notify/internal/logic"
	"cloudflare-pages-hook/notify/internal/svc"
	"cloudflare-pages-hook/notify/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func notifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.NotifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewNotifyLogic(r.Context(), svcCtx)
		resp, err := l.Notify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
