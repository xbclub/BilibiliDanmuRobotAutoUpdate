package main

import (
	"flag"
	"fmt"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/config"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/handler"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/autoupdate.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	err := ctx.Ghr.RefreshRelease(c.GithubInfo.User, c.GithubInfo.Repo)
	if err != nil {
		panic(err)
	}
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
