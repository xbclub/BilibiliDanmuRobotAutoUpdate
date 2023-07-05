package svc

import (
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/config"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/utiles"
)

type ServiceContext struct {
	Config config.Config
	Ghr    *utiles.ReleaseData
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Ghr:    utiles.NewGetRelease(),
	}
}
