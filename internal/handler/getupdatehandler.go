package handler

import (
	"net/http"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/logic"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetUpdateLogic(r.Context(), svcCtx)
		resp, err := l.GetUpdate()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
