## 前言

​	本项目是[xbclub/BilibiliDanmuRobot](https://github.com/xbclub/BilibiliDanmuRobot)的更新服务api

## 构建

Go： 1.20+

```yaml
go build autoupdate.go
```

## 接口

获取最新版本号及下载地址

| Url      | /getUpdate                                                   |
| -------- | ------------------------------------------------------------ |
| Method   | Get                                                          |
| Response | {"version":"v1.1.6","link":"https://ghproxy.com/https://github.com/xbclub/BilibiliDanmuRobot/releases/download/v1.1.6/GUI-BilibiliDanmuRobot_Windows_amd64_v1.1.6.zip","changeLog":"这里是markdown语法的更新日志"} |

