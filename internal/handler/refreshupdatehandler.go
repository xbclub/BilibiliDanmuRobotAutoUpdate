package handler

import (
	"net/http"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/logic"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func refreshUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RefreshUpdateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewRefreshUpdateLogic(r.Context(), svcCtx)
		resp, err := l.RefreshUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
