package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	GithubInfo struct {
		User string `json:",default=xbclub"`
		Repo string `json:",default=BilibiliDanmuRobot"`
	}
	UpgraderInfo struct {
		User string `json:",default=xbclub"`
		Repo string `json:",default=BilibiliDanmuRobot"`
	}
	Proxy  string `json:",default=https://github.moeyy.xyz/"`
	Secret string
}
