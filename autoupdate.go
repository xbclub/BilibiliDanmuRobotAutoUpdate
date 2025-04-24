package main

import (
	"flag"
	"fmt"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/config"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/handler"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
)

var configFile = flag.String("f", "etc/autoupdate.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	err := ctx.Ghr.RefreshRelease(c.GithubInfo.User, c.GithubInfo.Repo, c.Proxy)
	if err != nil {
		panic(err)
	}
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/test/*",
		Handler: http.FileServer(http.Dir("test")).ServeHTTP,
	})
	err = ctx.Upgrader.RefreshUpgraderRelease(c.UpgraderInfo.User, c.UpgraderInfo.Repo, c.Proxy)
	if err != nil {
		panic(err)
	}
	handler.RegisterHandlers(server, ctx)
	//RegisterHandlers(server)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

//func RegisterHandlers(engine *rest.Server) {
//
//	//这里注册
//	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
//	patern := "/test/"
//	dirpath := "./test/"
//	for i := 1; i < len(dirlevel); i++ {
//		path := patern + strings.Join(dirlevel[:i], "/")
//		//最后生成 /asset
//		engine.AddRoute(
//			rest.Route{
//				Method:  http.MethodGet,
//				Path:    path,
//				Handler: http.StripPrefix(patern, http.FileServer(http.Dir(dirpath))).ServeHTTP,
//			})
//
//		logx.Infof("register dir  %s  %s", path, dirpath)
//	}
//}
