### 项目介绍
~~~shell
针对于告警分组，把不同业务的告警信息发送至指定钉钉群,更多帮助查看doc/README.md
~~~
#### Windows 平台构建 Linux下运行
~~~shell
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build  -ldflags "-s -w" -o dingtalk  main.go
~~~


