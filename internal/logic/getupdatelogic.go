package logic

import (
	"context"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUpdateLogic {
	return &GetUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUpdateLogic) GetUpdate() (resp *types.GetUpdateResponse, err error) {
	return &types.GetUpdateResponse{
		Version:   l.svcCtx.Ghr.Version,
		Link:      l.svcCtx.Ghr.Link,
		ChangeLog: l.svcCtx.Ghr.Changelog,
	}, nil
}
