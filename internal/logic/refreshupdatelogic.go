package logic

import (
	"context"

	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/svc"
	"github.com/xbclub/BilibiliDanmuRobotAutoUpdate.git/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefreshUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRefreshUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefreshUpdateLogic {
	return &RefreshUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RefreshUpdateLogic) RefreshUpdate(req *types.RefreshUpdateReq) (resp *types.RefreshUpdateResp, err error) {
	if req.Secret != l.svcCtx.Config.Secret {
		return &types.RefreshUpdateResp{Code: 30100}, nil
	}
	err = l.svcCtx.Ghr.RefreshRelease(l.svcCtx.Config.GithubInfo.User, l.svcCtx.Config.GithubInfo.Repo)
	if err != nil {
		return &types.RefreshUpdateResp{Code: 50000}, err
	}
	return &types.RefreshUpdateResp{Code: 20000}, nil
}
