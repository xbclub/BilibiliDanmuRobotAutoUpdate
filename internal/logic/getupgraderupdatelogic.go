package logic

import (
	"context"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpgraderUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpgraderUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpgraderUpdateLogic {
	return &GetUpgraderUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpgraderUpdateLogic) GetUpgraderUpdate() (resp *types.GetUpdateResponse, err error) {
	logx.Info(l.svcCtx.Upgrader)
	return &types.GetUpdateResponse{
		Version:   l.svcCtx.Upgrader.Version,
		Link:      l.svcCtx.Upgrader.Link,
		ChangeLog: l.svcCtx.Upgrader.Changelog,
	}, nil
}
