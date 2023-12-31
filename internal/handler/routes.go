// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/getUpdate",
				Handler: getUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/refreshUpdate",
				Handler: refreshUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/getUpgraderUpdate",
				Handler: getUpgraderUpdateHandler(serverCtx),
			},
		},
	)
}
